package auth

import (
	"fmt"
	"kefu_go_robot/conf"
	"kefu_go_robot/utils"
)

// Token global auth token
var Token string

// FetchToken auth
func FetchToken() {
	config := new(conf.Cionfigs).GetConfigs()
	path := config.GatewayHost + "/v1/auth/token/"
	var request = map[string]string{}
	request["app_id"] = config.MiAppID
	request["app_key"] = config.MiAppKey
	request["app_secret"] = config.MiAppSecret
	response := utils.HTTPRequest(path, "POST", request, "")
	if response.Code != 200 {
		fmt.Println("/v1/auth/token/：授权出错！")
		return
	}
	Token = response.Data.(string)
}
