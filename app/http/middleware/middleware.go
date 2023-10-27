package middleware

import (
	"github.com/kataras/iris/v12"
	"ms/app/service/jwt"
	"strings"
)

func JwtAuthCheck(ctx iris.Context) {
	authorization := ctx.GetHeader("Authorization")
	if authorization == "" {
		ctx.StopWithJSON(401, iris.Map{"code": 401, "message": "token不能为空"})
		return
	}
	authArr := strings.Split(authorization, " ")
	if len(authArr) != 2 {
		ctx.StopWithJSON(401, iris.Map{"code": 401, "message": "token格式不正确，请使用Bearer token格式"})
		return
	}
	tokenString := authArr[1]
	claims, err := jwt.NewService().Check(tokenString)
	if err != nil {
		ctx.StopWithJSON(401, iris.Map{"code": 401, "message": err.Error()})
		return
	}
	ctx.Values().Set("user_id", claims.UserId)
	ctx.Values().Set("jwt_token", tokenString)
	ctx.Next()
}
