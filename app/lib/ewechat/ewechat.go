package ewechat

import (
	"github.com/ArtisanCloud/PowerWeChat/v3/src/work"
	"github.com/spf13/viper"
	"ms/app/model/resp"
)

func newClient() *work.Work {
	weComContactApp, _ := work.NewWork(&work.UserConfig{
		CorpID:  viper.GetString("enterpriseWechat.corp_id"), // 企业微信的app id，所有企业微信共用一个。
		AgentID: viper.GetInt("enterpriseWechat.agent_id"),   // 内部应用的app id
		Secret:  viper.GetString("enterpriseWechat.secret"),  // 内部应用的app secret
		//Token:       "[token]",        // 内部应用的app token
		//AESKey:      "[aes_key]",      // 内部应用的app aeskey
		//CallbackURL: "[callback_url]", // 内部应用的场景回调设置
		//OAuth: work.OAuth{
		//	Callback: "[app_oauth_callback_url]", // 内部应用的app oauth url
		//	Scopes:   nil,
		//},
		HttpDebug: true,
	})

	return weComContactApp
}

func GetUserInfo(code string) resp.SocialUserInfo {
	weComClient := newClient()
	user, err := weComClient.OAuth.Provider.GetUserInfo(code)
	if err != nil {
		panic(err)
	}
	userDetail, err := weComClient.OAuth.Provider.GetUserDetail(user.UserTicket)
	if err != nil {
		panic(err)
	}
	res := resp.SocialUserInfo{
		AvatarUrl: userDetail.Avatar,
		Email:     userDetail.Email,
		Mobile:    userDetail.Mobile,
		OpenId:    user.OpenID,
		UnionId:   user.UserID,
	}
	return res
}
