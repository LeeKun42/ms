package controller

import (
	"github.com/kataras/iris/v12"
	"ms/app/http/response"
	"ms/app/model/request"
	"ms/app/service/admin_user"
	"ms/app/service/jwt"
	"strings"
)

type AdminUserController struct {
	AdminUserService *admin_user.Service
	JwtService       *jwt.Service
}

func NewAdminUserController() *AdminUserController {
	return &AdminUserController{
		AdminUserService: admin_user.NewService(),
		JwtService:       jwt.NewService(),
	}
}

func (auc *AdminUserController) Lists(ctx iris.Context) {
	var params request.SearchAdminUserParams
	ctx.ReadQuery(&params)
	res := auc.AdminUserService.Lists(params)
	response.Success(ctx, res)
}

func (auc *AdminUserController) Create(ctx iris.Context) {
	// 获取请求参数
	var params request.CreateAdminUserParams
	ctx.ReadJSON(&params)
	// 检查参数
	if params.Mobile == "" || params.Email == "" || params.Name == "" || params.Passwd == "" {
		response.Fail(ctx, "参数不能为空")
		return
	}
	if strings.Count(params.Passwd, "") < 6 {
		response.Fail(ctx, "密码至少6位字符")
		return
	}

	adminUserId, err := auc.AdminUserService.Create(params)
	if err != nil {
		response.Fail(ctx, err.Error())
	} else {
		response.Success(ctx, iris.Map{"admin_user_id": adminUserId})
	}
}

func (auc *AdminUserController) Update(ctx iris.Context) {
	// 获取请求参数
	var params request.UpdateAdminUserParams
	ctx.ReadJSON(&params)
	adminUserId, _ := ctx.Params().GetInt("adminUserId")
	params.ID = adminUserId
	// 检查参数
	if params.Mobile == "" || params.Email == "" || params.Name == "" {
		response.Fail(ctx, "参数不能为空")
		return
	}

	err := auc.AdminUserService.Update(params)
	if err != nil {
		response.Fail(ctx, err.Error())
	} else {
		response.Success(ctx, iris.Map{"admin_user_id": adminUserId})
	}
}

func (auc *AdminUserController) Login(ctx iris.Context) {
	// 获取请求参数
	var params request.AdminUserLoginParams
	ctx.ReadJSON(&params)
	// 检查参数
	if params.Account == "" || params.Passwd == "" {
		response.Fail(ctx, "参数不能为空")
		return
	}
	token, err := auc.AdminUserService.Login(params)
	if err != nil {
		response.Fail(ctx, err.Error())
	} else {
		response.Success(ctx, iris.Map{"token": token})
	}
}

func (auc *AdminUserController) GetLoginUserInfo(ctx iris.Context) {
	// 获取请求参数
	loginUserId, _ := ctx.Values().GetInt("user_id")
	if loginUserId == 0 {
		response.Fail(ctx, "没有获取到登录用户id")
		return
	}
	adminUser := auc.AdminUserService.Info(loginUserId)
	response.Success(ctx, adminUser)
}

func (auc *AdminUserController) Logout(ctx iris.Context) {
	auc.JwtService.Invalidate(ctx.Values().GetString("jwt_token"))
	response.Success(ctx, iris.Map{})
}
