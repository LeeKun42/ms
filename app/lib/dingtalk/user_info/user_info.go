package user_info

import (
	"fmt"
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	dingtalkcontact_1_0 "github.com/alibabacloud-go/dingtalk/contact_1_0"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
	"ms/app/lib/dingtalk/user_token"
	"ms/app/model/resp"
)

func newClient() *dingtalkcontact_1_0.Client {
	ac := &openapi.Config{
		Protocol: tea.String("https"),
		RegionId: tea.String("central"),
	}
	client, _ := dingtalkcontact_1_0.NewClient(ac)
	return client
}

func GetUserInfo(code string) resp.SocialUserInfo {
	getUserHeaders := &dingtalkcontact_1_0.GetUserHeaders{}
	getUserHeaders.XAcsDingtalkAccessToken = tea.String(user_token.GetAccessToken(code))
	res, err := newClient().GetUserWithOptions(tea.String("me"), getUserHeaders, &util.RuntimeOptions{})
	var user resp.SocialUserInfo
	if err != nil {
		fmt.Println(err)
	} else {
		user.Email = *res.Body.Email
		user.AvatarUrl = *res.Body.AvatarUrl
		user.Mobile = *res.Body.Mobile
		user.OpenId = *res.Body.OpenId
		user.UnionId = *res.Body.UnionId
	}
	return user
}
