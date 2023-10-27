package http

import (
	"context"
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/cors"
	recover2 "github.com/kataras/iris/v12/middleware/recover"
	"github.com/spf13/viper"
	"ms/app/http/controller"
	"ms/app/http/middleware"
	"ms/app/lib/log"
	"time"
)

// StartWebServer 开启web服务
func StartWebServer() {
	app := iris.New()
	app.Use(recover2.New())
	//跨域中间件
	app.UseRouter(cors.New().Handler())
	app.Use(log.HttpLogHandler)

	iris.RegisterOnInterrupt(func() {
		timeout := 5 * time.Second
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()
		fmt.Println("Shutdown server")
		// 关闭所有主机
		app.Shutdown(ctx)
	})

	/** restful风格路由 */
	app.PartyFunc("/api", func(api iris.Party) {
		adminUserController := controller.NewAdminUserController()
		api.PartyFunc("/admin-users", func(user iris.Party) {
			user.Post("/token", adminUserController.Login)
			user.Get("/0", adminUserController.GetLoginUserInfo).Use(middleware.JwtAuthCheck)
			user.Post("", adminUserController.Create).Use(middleware.JwtAuthCheck)
			user.Delete("/token", adminUserController.Logout).Use(middleware.JwtAuthCheck)
			user.Get("", adminUserController.Lists).Use(middleware.JwtAuthCheck)
			user.Patch("/{adminUserId:int}", adminUserController.Update).Use(middleware.JwtAuthCheck)

			socialController := controller.NewAdminUserSocialController()
			user.Get("/social/{code:string}", socialController.LoginByCode)
			user.Get("/social", socialController.GetSocialAccount).Use(middleware.JwtAuthCheck)
			user.Post("/social/{code:string}", socialController.BindSocialAccount).Use(middleware.JwtAuthCheck)
			user.Delete("/social/{type:string}", socialController.UnBindSocialAccount).Use(middleware.JwtAuthCheck)
		})

		roleController := controller.NewRoleController()
		api.PartyFunc("/roles", func(role iris.Party) {
			role.Post("", roleController.Create).Use(middleware.JwtAuthCheck)
			role.Get("", roleController.Lists).Use(middleware.JwtAuthCheck)
			role.Get("/{roleId:int}", roleController.Info).Use(middleware.JwtAuthCheck)
			role.Get("/option", roleController.AllOptions).Use(middleware.JwtAuthCheck)
			role.Patch("/{roleId:int}", roleController.Update).Use(middleware.JwtAuthCheck)
			role.Delete("/{roleId:int}", roleController.Delete).Use(middleware.JwtAuthCheck)
			role.Patch("/{roleId:int}/permissions", roleController.AttachPermissions).Use(middleware.JwtAuthCheck)
		})

		permissionController := controller.NewPermissionController()
		api.PartyFunc("/permissions", func(permission iris.Party) {
			permission.Post("", permissionController.Create).Use(middleware.JwtAuthCheck)
			permission.Get("", permissionController.Lists).Use(middleware.JwtAuthCheck)
			permission.Get("/option", permissionController.AllOptions).Use(middleware.JwtAuthCheck)
			permission.Get("/parent", permissionController.ParentOptions).Use(middleware.JwtAuthCheck)
			permission.Patch("/{permissionId:int}", permissionController.Update).Use(middleware.JwtAuthCheck)
			permission.Delete("/{permissionId:int}", permissionController.Delete).Use(middleware.JwtAuthCheck)
		})
	})

	port := viper.GetInt("server.http")

	/**
	开启web服务
	参数1：监听地址和端口
	参数2：允许body多次消费
	*/
	app.Run(iris.Addr(fmt.Sprintf(":%d", port)), iris.WithoutBodyConsumptionOnUnmarshal)
}
