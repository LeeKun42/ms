package permission

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
func (service *Service) Lists(params request.SearchPermissionParams) (permissionListRes resp.PermissionListResponse) {
	permissionListRes.Meta.PageIndex = params.PageIndex
	permissionListRes.Meta.PageSize = params.PageSize
	offset := (params.PageIndex - 1) * params.PageSize
	tx := service.DB.Table(dto.TableNamePermission)
	if params.ParentID != -1 {
		tx.Where("parent_id", params.ParentID)
	}
	if params.Flag != "" {
		tx.Where("flag like ?", "%"+params.Flag+"%")
	}
	if params.Name != "" {
		tx.Where("name like ?", "%"+params.Name+"%")
	}
	if params.Type != 0 {
		tx.Where("type", params.Type)
	}
	tx.Count(&permissionListRes.Meta.Total)

	var permissions []dto.Permission
	if params.PageSize == 0 && params.PageIndex == 0 {
		permissionListRes.Meta.PageCount = 1
		tx.Order("id ASC")
	} else {
		permissionListRes.Meta.PageCount = (int(permissionListRes.Meta.Total-1) / permissionListRes.Meta.PageSize) + 1
		tx.Order("id DESC").Offset(offset).Limit(params.PageSize)
	}
	tx.Find(&permissions)
	for _, permission := range permissions {
		permissionListRes.Data = append(permissionListRes.Data, permission.ToResponse())
	}
	return
}

func (service *Service) TreeList(list []resp.PermissionResponse) (res []resp.PermissionResponse) {
	for i := 0; i < len(list); i++ {
		permissionResponse := list[i]
		if permissionResponse.ParentID == 0 {
			for k := 0; k < len(list); k++ {
				permission := list[k]
				if permissionResponse.ID == permission.ParentID {
					permissionResponse.Children = append(permissionResponse.Children, permission)
				}
			}
			res = append(res, permissionResponse)
		}
	}
	return
}

// Create 创建权限
func (service *Service) Create(params request.CreatePermissionParams) (int, error) {
	var permissionDTO dto.Permission
	service.DB.Table(permissionDTO.TableName()).Where("flag", params.Flag).First(&permissionDTO)
	if permissionDTO.ID != 0 {
		return 0, errors.New("创建失败：权限标识已存在")
	}
	permissionDTO = dto.Permission{
		ParentID: params.ParentID,
		Flag:     params.Flag,
		Name:     params.Name,
		Desc:     params.Desc,
		Type:     params.Type,
	}
	result := service.DB.Create(&permissionDTO)
	if result.Error != nil {
		return 0, errors.New("创建失败：" + result.Error.Error())
	}
	return permissionDTO.ID, nil
}

// Update 更新权限
func (service *Service) Update(params request.CreatePermissionParams) error {
	var permissionDTO dto.Permission
	service.DB.Table(permissionDTO.TableName()).Where("id", params.ID).First(&permissionDTO)
	if permissionDTO.ID == 0 {
		return errors.New("更新失败：权限不存在")
	}
	tx := service.DB.Table(permissionDTO.TableName()).Where("id", params.ID).
		Updates(dto.Permission{
			ParentID: params.ParentID,
			Flag:     params.Flag,
			Name:     params.Name,
			Desc:     params.Desc,
			Type:     params.Type,
		})
	return tx.Error
}

// Delete 删除权限
func (service *Service) Delete(permissionID int) error {
	tx := service.DB.Begin()
	// 删除角色关联权限数据
	err := tx.Where("permission_id", permissionID).Delete(&dto.RolePermission{}).Error
	if err != nil {
		tx.Rollback()
	}
	// 删除权限数据
	err = tx.Where("id", permissionID).Delete(&dto.Permission{}).Error
	if err != nil {
		tx.Rollback()
	}
	tx.Commit()
	return nil
}
