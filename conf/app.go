package conf

// see https://admin.mimc.chat.xiaomi.net/docs/02-createapp.html

// AppConfigsObject ...
type AppConfigsObject struct {
	httpURL    string // 小米消息云接口地址
	AppID      int64  // AppID
	AppKey     string // AppKey
	AppSecret  string // AppSecret
	AppAccount string // AppAccount
}

// GetAppConfigs ...
func GetAppConfigs() AppConfigsObject {
	return AppConfigsObject{
		httpURL:    "",
		AppID:      "",
		AppKey:     "",
		AppSecret:  "",
		AppAccount: "",
	}
}
