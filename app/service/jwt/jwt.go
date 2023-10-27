package jwt

import (
	"context"
	"errors"
	jwtv5 "github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
	"ms/app/lib/redis"
	"strconv"
	"time"
)

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

type CustomClaims struct {
	UserId int `json:"user_id"`
	jwtv5.RegisteredClaims
}

const UserTypeClient = "client"
const UserTypeAdmin = "admin"

var redisKeyPrefixMap = map[string]string{
	"client": "client:jwt:token:",
	"admin":  "admin:jwt:token:",
}

func (js *Service) Create(userId int, userType string) string {
	//生成token
	now := time.Now()              //当前时间
	ttl := viper.GetInt("jwt.ttl") //token有效期（分钟）

	exp := now.Add(time.Minute * time.Duration(ttl)) //过期时间
	//自定义jwt body内容
	claims := CustomClaims{
		userId,
		jwtv5.RegisteredClaims{
			Issuer:    "ms",
			Subject:   userType,
			Audience:  nil,
			ExpiresAt: jwtv5.NewNumericDate(exp),
			NotBefore: jwtv5.NewNumericDate(now),
			IssuedAt:  jwtv5.NewNumericDate(now),
			ID:        strconv.Itoa(userId),
		},
	}
	//生成jwt签名
	tk := jwtv5.NewWithClaims(jwtv5.SigningMethodHS256, claims)
	token, _ := tk.SignedString([]byte(viper.GetString("jwt.secret")))

	//写入redis 用以实现指定token失效功能
	ctx, _ := context.WithTimeout(context.Background(), time.Millisecond*500)
	key := redisKeyPrefixMap[userType] + token
	redis.Cache().Set(ctx, key, exp.Unix(), time.Minute*time.Duration(ttl))

	return token
}

func (js *Service) Check(tokenString string) (*CustomClaims, error) {
	token, err := jwtv5.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwtv5.Token) (interface{}, error) {
		return []byte(viper.GetString("jwt.secret")), nil
	})
	if err != nil {
		return &CustomClaims{}, errors.New(err.Error())
	}
	claims, ok := token.Claims.(*CustomClaims)
	if ok && token.Valid {
		//检查是否在redis中存在
		ctx, _ := context.WithTimeout(context.Background(), time.Millisecond*500)
		userType := claims.Subject
		key := redisKeyPrefixMap[userType] + tokenString
		_, err := redis.Cache().Get(ctx, key).Result()
		if err != nil { //
			return &CustomClaims{}, errors.New("无效的token")
		} else {
			//todo 检查密码是否已修改
			return claims, nil
		}
	} else {
		return &CustomClaims{}, errors.New("token反序列化错误")
	}
}

func (js *Service) Invalidate(tokenString string) error {
	token, err := jwtv5.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwtv5.Token) (interface{}, error) {
		return []byte(viper.GetString("jwt.secret")), nil
	})
	if err != nil {
		return errors.New("token无效")
	}
	claims, ok := token.Claims.(*CustomClaims)
	if ok && token.Valid {
		ctx, _ := context.WithTimeout(context.Background(), time.Millisecond*500)
		userType := claims.Subject
		key := redisKeyPrefixMap[userType] + tokenString
		return redis.Cache().Del(ctx, key).Err()
	} else {
		return errors.New("token无效")
	}
}
