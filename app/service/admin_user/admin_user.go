package admin_user

import (
	"errors"
	"gorm.io/gorm"
	"ms/app/lib/hash"
	"ms/app/model"
	"ms/app/model/dto"
	"ms/app/model/request"
	"ms/app/model/resp"
	"ms/app/service/jwt"
)

type Service struct {
	DB *gorm.DB
}

func NewService() *Service {
	return &Service{
		DB: model.Instance(),
	}
}

// Create 创建管理账号
func (service *Service) Create(params request.CreateAdminUserParams) (int, error) {
	var adminUserDTO dto.AdminUser
	service.DB.Where("mobile", params.Mobile).Or("email", params.Email).First(&adminUserDTO)
	if adminUserDTO.ID != 0 {
		return 0, errors.New("创建失败：手机号或者邮箱已存在")
	}
	adminUserDTO = dto.AdminUser{
		Mobile: params.Mobile,
		Email:  params.Email,
		Name:   params.Name,
		Passwd: hash.Make(params.Passwd),
		Status: params.Status,
	}
	tx := service.DB.Begin()
	result := tx.Create(&adminUserDTO)
	if result.Error != nil {
		tx.Rollback()
		return 0, errors.New("创建失败：" + result.Error.Error())
	}
	err := service.attachRoles(tx, adminUserDTO.ID, params.Roles)
	if err != nil {
		tx.Rollback()
		return 0, errors.New("创建失败：" + err.Error())
	}
	tx.Commit()
	return adminUserDTO.ID, nil
}

// Update 更新管理账号
func (service *Service) Update(params request.UpdateAdminUserParams) error {
	var adminUserDTO dto.AdminUser
	service.DB.Where("id", params.ID).First(&adminUserDTO)
	if adminUserDTO.ID == 0 {
		return errors.New("更新失败：账号不存在")
	}
	var cnt int64
	service.DB.Table(adminUserDTO.TableName()).Where("(mobile=? or email=?)", params.Mobile, params.Email).Where("id <> ?", params.ID).Count(&cnt)
	if cnt > 0 {
		return errors.New("更新失败：手机号或者邮箱已存在")
	}
	adminUserDTO.Mobile = params.Mobile
	adminUserDTO.Email = params.Email
	adminUserDTO.Name = params.Name
	adminUserDTO.Status = params.Status
	tx := service.DB.Begin()
	result := tx.Save(&adminUserDTO)
	if result.Error != nil {
		tx.Rollback()
		return errors.New("更新失败：" + result.Error.Error())
	}
	err := service.attachRoles(tx, adminUserDTO.ID, params.Roles)
	if err != nil {
		tx.Rollback()
		return errors.New("创建失败：" + err.Error())
	}
	tx.Commit()
	return nil
}

func (service *Service) Login(params request.AdminUserLoginParams) (string, error) {
	var adminUserDTO dto.AdminUser
	result := service.DB.Where("mobile", params.Account).Or("email", params.Account).First(&adminUserDTO)
	if result.Error != nil && errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return "", errors.New("账号不存在")
	}
	if adminUserDTO.Status == dto.UserStatusDisabled {
		return "", errors.New("账号已被禁用，请联系管理员")
	}
	if !hash.Check(params.Passwd, adminUserDTO.Passwd) { //密码不正确
		return "", errors.New("密码不正确")
	}
	token := jwt.NewService().Create(adminUserDTO.ID, jwt.UserTypeAdmin)
	return token, nil
}

// Lists 查询管理员账号列表
func (service *Service) Lists(params request.SearchAdminUserParams) (adminUserListRes resp.AdminUserListResponse) {
	adminUserListRes.Meta.PageIndex = params.PageIndex
	adminUserListRes.Meta.PageSize = params.PageSize
	offset := (params.PageIndex - 1) * params.PageSize
	tx := service.DB.Table(dto.TableNameAdminUser)
	if params.Mobile != "" {
		tx.Where("mobile like ?", "%"+params.Mobile+"%")
	}
	if params.Email != "" {
		tx.Where("email like ?", "%"+params.Email+"%")
	}
	if params.Name != "" {
		tx.Where("name like ?", "%"+params.Name+"%")
	}
	if params.Status != -1 {
		tx.Where("status", params.Status)
	}
	if params.RoleId > 0 {
		tx.Where("EXISTS(SELECT b.admin_user_id FROM admin_user_roles as b WHERE b.role_id=? and b.admin_user_id=admin_users.id)", params.RoleId)
	}
	tx.Count(&adminUserListRes.Meta.Total)
	adminUserListRes.Meta.PageCount = (int(adminUserListRes.Meta.Total-1) / adminUserListRes.Meta.PageSize) + 1
	var adminUserList []dto.AdminUser
	tx.Order("id DESC").Preload("Roles").Offset(offset).Limit(params.PageSize).Find(&adminUserList)
	for _, adminUser := range adminUserList {
		adminUserListRes.Data = append(adminUserListRes.Data, adminUser.ToResponse())
	}
	return
}

func (service *Service) Info(adminUserId int) resp.AdminUserResponse {
	var adminUserDTO dto.AdminUser
	service.DB.Where("id", adminUserId).Preload("Roles.Permissions").First(&adminUserDTO)
	return adminUserDTO.ToResponse()
}

// AttachRoles 给角色关联权限
func (service *Service) attachRoles(tx *gorm.DB, adminUserID int, roleIds []int) error {
	// 删除原有权限
	err := tx.Where("admin_user_id", adminUserID).Delete(&dto.AdminUserRole{}).Error
	if err != nil {
		return err
	}
	// 关联新选择的权限
	for _, roleId := range roleIds {
		err = tx.Create(dto.AdminUserRole{AdminUserID: adminUserID, RoleID: roleId}).Error
		if err != nil {
			return err
		}
	}
	return nil
}
