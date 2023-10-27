package admin_user_social

import (
	"errors"
	"gorm.io/gorm"
	"ms/app/lib/dingtalk/user_info"
	"ms/app/lib/ewechat"
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

const SocialType_Wechat = "wechat"
const SocialType_Dingding = "dingtalk"
const SocialType_Enterprisewechat = "enterprisewechat"

var SocialTypeMap = map[string]string{
	SocialType_Wechat:           "微信",
	SocialType_Dingding:         "钉钉",
	SocialType_Enterprisewechat: "企业微信",
}

func (service *Service) CodeToSocialUser(code string, socialType string) (res resp.SocialUserInfo) {
	switch socialType {
	case SocialType_Dingding:
		res = user_info.GetUserInfo(code)
	case SocialType_Enterprisewechat:
		res = ewechat.GetUserInfo(code)
	}
	return
}

func (service *Service) Get(adminUserID int, socialType string) dto.AdminUserSocial {
	var social dto.AdminUserSocial
	service.DB.Where("admin_user_id", adminUserID).Where("type", socialType).First(&social)
	return social
}

func (service *Service) Create(params request.BindSocialParams) (dto.AdminUserSocial, error) {
	var adminUserSocial dto.AdminUserSocial
	service.DB.Where("type", params.Type).Where("admin_user_id", params.AdminUserID).First(&adminUserSocial)
	if adminUserSocial.AdminUserID != 0 {
		return adminUserSocial, nil
	}

	adminUserSocial = dto.AdminUserSocial{
		AdminUserID: params.AdminUserID,
		Type:        params.Type,
		Unionid:     params.Unionid,
		Openid:      params.Openid,
	}
	result := service.DB.Create(&adminUserSocial)
	if result.Error != nil {
		return adminUserSocial, errors.New("关联失败：" + result.Error.Error())
	}
	var adminUserDTO dto.AdminUser
	service.DB.Where("id", adminUserSocial.AdminUserID).First(&adminUserDTO)
	if adminUserDTO.Avatar == "" {
		service.DB.Model(&adminUserDTO).Where("id", adminUserSocial.AdminUserID).Update("avatar", params.Avatar)
	}
	return adminUserSocial, nil
}

func (service *Service) Delete(adminUserID int, socialType string) error {
	return service.DB.Where("type", socialType).Where("admin_user_id", adminUserID).Delete(&dto.AdminUserSocial{}).Error
}

func (service *Service) Login(unionID string, socialType string) (string, error) {
	var adminUserSocial dto.AdminUserSocial
	service.DB.Where("type", socialType).Where("unionid", unionID).First(&adminUserSocial)
	if adminUserSocial.AdminUserID == 0 {
		return "", errors.New("未关联的" + SocialTypeMap[socialType] + "账号，无法登录")
	}
	var adminUserDTO dto.AdminUser
	service.DB.Where("id", adminUserSocial.AdminUserID).First(&adminUserDTO)
	if adminUserDTO.ID == 0 {
		return "", errors.New("账号不存在！")
	}
	token := jwt.NewService().Create(adminUserDTO.ID, jwt.UserTypeAdmin)
	return token, nil
}
