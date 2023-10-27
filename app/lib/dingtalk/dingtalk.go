package dingtalk

import "github.com/spf13/viper"

type config struct {
	AppID     string `mapstructure:"key"`
	AppSecret string `mapstructure:"secret"`
}

func GetAppID() string {
	return viper.GetString("dingTalk.key")
}

func GetAppSecret() string {
	return viper.GetString("dingTalk.secret")
}
