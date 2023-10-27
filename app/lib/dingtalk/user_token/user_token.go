package user_token

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	dingtalkoauth2_1_0 "github.com/alibabacloud-go/dingtalk/oauth2_1_0"
	"github.com/alibabacloud-go/tea/tea"
	"ms/app/lib/dingtalk"
)

func newClient() *dingtalkoauth2_1_0.Client {
	ac := &openapi.Config{
		Protocol: tea.String("https"),
		RegionId: tea.String("central"),
	}
	client, _ := dingtalkoauth2_1_0.NewClient(ac)
	return client
}

func GetAccessToken(code string) string {
	getUserTokenRequest := &dingtalkoauth2_1_0.GetUserTokenRequest{
		ClientId:     tea.String(dingtalk.GetAppID()),
		ClientSecret: tea.String(dingtalk.GetAppSecret()),
		Code:         tea.String(code),
		RefreshToken: tea.String(""),
		GrantType:    tea.String("authorization_code"),
	}
	res, err := newClient().GetUserToken(getUserTokenRequest)
	if err != nil {
		return ""
	}
	return *res.Body.AccessToken
}
