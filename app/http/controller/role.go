package controller

import (
	"github.com/kataras/iris/v12"
	"ms/app/http/response"
	"ms/app/model/request"
	"ms/app/service/role"
)

type RoleController struct {
	RoleService *role.Service
}

func NewRoleController() *RoleController {
	return &RoleController{
		RoleService: role.NewService(),
	}
}

func (rc *RoleController) Lists(ctx iris.Context) {
	// 获取请求参数
	var params request.SearchRoleParams
	ctx.ReadQuery(&params)
	res := rc.RoleService.Lists(params)
	response.Success(ctx, res)
}

func (rc *RoleController) AllOptions(ctx iris.Context) {
	// 获取请求参数
	params := request.SearchRoleParams{}
	params.PageIndex = 0
	params.PageSize = 0
	res := rc.RoleService.Lists(params)
	response.Success(ctx, res.Data)
}

func (rc *RoleController) Create(ctx iris.Context) {
	// 获取请求参数
	var params request.CreateRoleParams
	ctx.ReadJSON(&params)
	// 检查参数
	if params.Flag == "" || params.Name == "" {
		response.Fail(ctx, "参数不能为空")
		return
	}

	roleId, err := rc.RoleService.Create(params)
	if err != nil {
		response.Fail(ctx, err.Error())
	} else {
		response.Success(ctx, iris.Map{"role_id": roleId})
	}
}

func (rc *RoleController) Info(ctx iris.Context) {
	// 获取请求参数
	roleId, _ := ctx.Params().GetInt("roleId")
	res := rc.RoleService.Info(roleId)
	response.Success(ctx, res)
}

func (rc *RoleController) Update(ctx iris.Context) {
	// 获取请求参数
	var params request.CreateRoleParams
	ctx.ReadJSON(&params)
	roleId, _ := ctx.Params().GetInt("roleId")
	params.ID = roleId
	// 检查参数
	if params.Flag == "" || params.Name == "" {
		response.Fail(ctx, "参数不能为空")
		return
	}

	err := rc.RoleService.Update(params)
	if err != nil {
		response.Fail(ctx, err.Error())
	} else {
		response.Success(ctx, iris.Map{"role_id": roleId})
	}
}

func (rc *RoleController) Delete(ctx iris.Context) {
	roleId, _ := ctx.Params().GetInt("roleId")
	// 检查参数
	if roleId == 0 {
		response.Fail(ctx, "参数不能为空")
		return
	}

	err := rc.RoleService.Delete(roleId)
	if err != nil {
		response.Fail(ctx, err.Error())
	} else {
		response.Success(ctx, iris.Map{})
	}
}

func (rc *RoleController) AttachPermissions(ctx iris.Context) {
	// 获取请求参数
	var params request.SetRolePermissionsParams
	ctx.ReadJSON(&params)
	roleId, _ := ctx.Params().GetInt("roleId")
	params.RoleID = roleId
	err := rc.RoleService.AttachPermissions(params)
	if err != nil {
		response.Fail(ctx, err.Error())
	} else {
		response.Success(ctx, iris.Map{})
	}
}
