package role

import (
	"errors"
	"gorm.io/gorm"
	"ms/app/model"
	"ms/app/model/dto"
	"ms/app/model/request"
	"ms/app/model/resp"
)

type Service struct {
	DB *gorm.DB
}

func NewService() *Service {
	return &Service{
		DB: model.Instance(),
	}
}

// Lists 查询管理员账号列表
func (service *Service) Lists(params request.SearchRoleParams) (roleListRes resp.RoleListResponse) {
	roleListRes.Meta.PageIndex = params.PageIndex
	roleListRes.Meta.PageSize = params.PageSize
	offset := (params.PageIndex - 1) * params.PageSize
	tx := service.DB.Table(dto.TableNameRole)
	if params.Flag != "" {
		tx.Where("flag like ?", "%"+params.Flag+"%")
	}
	if params.Name != "" {
		tx.Where("name like ?", "%"+params.Name+"%")
	}
	tx.Count(&roleListRes.Meta.Total)

	var roles []dto.Role
	if params.PageSize == 0 && params.PageIndex == 0 {
		roleListRes.Meta.PageCount = 1
		tx.Order("id ASC")
	} else {
		roleListRes.Meta.PageCount = (int(roleListRes.Meta.Total-1) / roleListRes.Meta.PageSize) + 1
		tx.Order("id DESC").Offset(offset).Limit(params.PageSize)
	}
	tx.Find(&roles)
	for _, role := range roles {
		roleListRes.Data = append(roleListRes.Data, role.ToResponse())
	}
	return
}

func (service *Service) Info(id int) (roleRes resp.RoleResponse) {
	var roleDTO dto.Role
	service.DB.Where("id", id).Preload("Permissions").First(&roleDTO)
	roleRes = roleDTO.ToResponse()
	return
}

// Create 创建角色
func (service *Service) Create(params request.CreateRoleParams) (int, error) {
	var roleDTO dto.Role
	service.DB.Table(roleDTO.TableName()).Where("flag", params.Flag).First(&roleDTO)
	if roleDTO.ID != 0 {
		return 0, errors.New("创建失败：角色标识已存在")
	}
	roleDTO = dto.Role{
		Flag:     params.Flag,
		Name:     params.Name,
		IsSystem: 0,
	}
	result := service.DB.Create(&roleDTO)
	if result.Error != nil {
		return 0, errors.New("创建失败：" + result.Error.Error())
	}
	return roleDTO.ID, nil
}

// Update 更新角色
func (service *Service) Update(params request.CreateRoleParams) error {
	var roleDTO dto.Role
	service.DB.Table(roleDTO.TableName()).Where("id", params.ID).First(&roleDTO)
	if roleDTO.ID == 0 {
		return errors.New("更新失败：角色不存在")
	}
	tx := service.DB.Table(roleDTO.TableName()).Where("id", params.ID).
		Updates(dto.Role{Flag: params.Flag, Name: params.Name})
	return tx.Error
}

// Delete 删除角色
func (service *Service) Delete(roleID int) error {
	tx := service.DB.Begin()
	// 删除角色关联用户数据
	err := tx.Where("role_id", roleID).Delete(&dto.AdminUserRole{}).Error
	if err != nil {
		tx.Rollback()
	}
	// 删除角色关联权限数据
	err = tx.Where("role_id", roleID).Delete(&dto.RolePermission{}).Error
	if err != nil {
		tx.Rollback()
	}
	// 删除角色数据
	err = tx.Where("id", roleID).Delete(&dto.Role{}).Error
	if err != nil {
		tx.Rollback()
	}
	tx.Commit()
	return nil
}

// AttachPermissions 给角色关联权限
func (service *Service) AttachPermissions(params request.SetRolePermissionsParams) error {
	var roleDTO dto.Role
	service.DB.Table(roleDTO.TableName()).Where("id", params.RoleID).First(&roleDTO)
	if roleDTO.ID == 0 {
		return errors.New("关联权限失败：角色不存在")
	}
	tx := service.DB.Begin()
	// 删除原有权限
	err := tx.Where("role_id", params.RoleID).Delete(&dto.RolePermission{}).Error
	if err != nil {
		tx.Rollback()
	}
	// 关联新选择的权限
	for _, permissionId := range params.PermissionIds {
		err = tx.Create(dto.RolePermission{RoleID: params.RoleID, PermissionID: permissionId}).Error
		if err != nil {
			tx.Rollback()
		}
	}
	tx.Commit()
	return nil
}
