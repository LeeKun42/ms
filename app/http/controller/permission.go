package controller

import (
	"github.com/kataras/iris/v12"
	"ms/app/http/response"
	"ms/app/model/request"
	"ms/app/service/permission"
)

type PermissionController struct {
	PermissionService *permission.Service
}

func NewPermissionController() *PermissionController {
	return &PermissionController{
		PermissionService: permission.NewService(),
	}
}

func (pc *PermissionController) Lists(ctx iris.Context) {
	// 获取请求参数
	var params request.SearchPermissionParams
	ctx.ReadQuery(&params)
	res := pc.PermissionService.Lists(params)
	response.Success(ctx, res)
}

func (pc *PermissionController) ParentOptions(ctx iris.Context) {
	// 构造查询参数
	params := request.SearchPermissionParams{
		ParentID: 0,
		Type:     10,
	}
	params.PageIndex = 0
	params.PageSize = 0 //不分页
	res := pc.PermissionService.Lists(params)
	response.Success(ctx, res)
}

func (pc *PermissionController) AllOptions(ctx iris.Context) {
	// 构造查询参数
	params := request.SearchPermissionParams{
		ParentID: -1,
		Type:     0,
	}
	params.PageIndex = 0
	params.PageSize = 0
	list := pc.PermissionService.Lists(params)
	res := pc.PermissionService.TreeList(list.Data)
	response.Success(ctx, res)
}

func (pc *PermissionController) Create(ctx iris.Context) {
	// 获取请求参数
	var params request.CreatePermissionParams
	ctx.ReadJSON(&params)
	// 检查参数
	if params.Flag == "" || params.Name == "" {
		response.Fail(ctx, "参数不能为空")
		return
	}

	permissionId, err := pc.PermissionService.Create(params)
	if err != nil {
		response.Fail(ctx, err.Error())
	} else {
		response.Success(ctx, iris.Map{"permission_id": permissionId})
	}
}

func (pc *PermissionController) Update(ctx iris.Context) {
	// 获取请求参数
	var params request.CreatePermissionParams
	ctx.ReadJSON(&params)
	permissionId, _ := ctx.Params().GetInt("permissionId")
	params.ID = permissionId
	// 检查参数
	if params.Flag == "" || params.Name == "" {
		response.Fail(ctx, "参数不能为空")
		return
	}

	err := pc.PermissionService.Update(params)
	if err != nil {
		response.Fail(ctx, err.Error())
	} else {
		response.Success(ctx, iris.Map{"permission_id": permissionId})
	}
}

func (pc *PermissionController) Delete(ctx iris.Context) {
	permissionId, _ := ctx.Params().GetInt("permissionId")
	// 检查参数
	if permissionId == 0 {
		response.Fail(ctx, "参数不能为空")
		return
	}

	err := pc.PermissionService.Delete(permissionId)
	if err != nil {
		response.Fail(ctx, err.Error())
	} else {
		response.Success(ctx, iris.Map{})
	}
}
