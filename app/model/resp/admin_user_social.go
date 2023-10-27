package resp

type SocialUserInfo struct {
	AvatarUrl string `json:"avatarUrl,omitempty"`
	Email     string `json:"email,omitempty"`
	Mobile    string `json:"mobile,omitempty"`
	OpenId    string `json:"openId,omitempty"`
	UnionId   string `json:"unionId,omitempty"`
}
