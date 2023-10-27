package request

type BindSocialParams struct {
	AdminUserID int    `json:"admin_user_id"`
	Type        string `json:"type"`
	Unionid     string `json:"unionid"`
	Openid      string `json:"openid"`
	Avatar      string `json:"avatar"`
}

type SocialParams struct {
	Type string
}
