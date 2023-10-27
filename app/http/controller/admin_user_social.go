package controller

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"ms/app/http/response"
	"ms/app/model/request"
	"ms/app/service/admin_user_social"
	"ms/app/service/jwt"
)

type AdminUserSocialController struct {
	AdminUserSocialService *admin_user_social.Service
	JwtService             *jwt.Service
}

func NewAdminUserSocialController() *AdminUserSocialController {
	return &AdminUserSocialController{
		AdminUserSocialService: admin_user_social.NewService(),
		JwtService:             jwt.NewService(),
	}
}

func (dc *AdminUserSocialController) LoginByCode(ctx iris.Context) {
	code := ctx.Params().Get("code")
	var params request.SocialParams
	ctx.ReadQuery(&params)
	res := dc.AdminUserSocialService.CodeToSocialUser(code, params.Type)
	token, err := dc.AdminUserSocialService.Login(res.UnionId, params.Type)
	if err != nil {
		response.Fail(ctx, err.Error())
	} else {
		response.Success(ctx, iris.Map{"token": token})
	}
}

func (dc *AdminUserSocialController) BindSocialAccount(ctx iris.Context) {
	code := ctx.Params().Get("code")
	var params request.SocialParams
	ctx.ReadJSON(&params)
	fmt.Println("BindSocialAccount", params.Type)
	res := dc.AdminUserSocialService.CodeToSocialUser(code, params.Type)
	adminUserID, _ := ctx.Values().GetInt("user_id")
	socialParams := request.BindSocialParams{
		AdminUserID: adminUserID,
		Type:        params.Type,
		Unionid:     res.UnionId,
		Openid:      res.OpenId,
		Avatar:      res.AvatarUrl,
	}
	social, err := dc.AdminUserSocialService.Create(socialParams)
	if err != nil {
		response.Fail(ctx, err.Error())
	} else {
		response.Success(ctx, social)
	}
}

func (dc *AdminUserSocialController) UnBindSocialAccount(ctx iris.Context) {
	adminUserID, _ := ctx.Values().GetInt("user_id")
	socialType := ctx.Params().Get("type")
	err := dc.AdminUserSocialService.Delete(adminUserID, socialType)
	if err != nil {
		response.Fail(ctx, err.Error())
	} else {
		response.Success(ctx, iris.Map{})
	}
}

func (dc *AdminUserSocialController) GetSocialAccount(ctx iris.Context) {
	adminUserID, _ := ctx.Values().GetInt("user_id")
	var params request.SocialParams
	ctx.ReadQuery(&params)
	fmt.Println("GetSocialAccount", params.Type)
	res := dc.AdminUserSocialService.Get(adminUserID, params.Type)
	response.Success(ctx, res)
}
