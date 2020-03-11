package services

import (
	"fmt"
	"kefu_go_robot/conf"
	"kefu_server/utils"
)

// AuthToken token
var AuthToken string

// AuthTokenRepository struct
type AuthTokenRepository struct{}

// GetAuthTokenRepositoryInstance get instance
func GetAuthTokenRepositoryInstance() *AuthTokenRepository {
	instance := new(AuthTokenRepository)
	return instance
}

// FetchToken auth
func (r *AuthTokenRepository) FetchToken() {
	config := new(conf.Cionfigs).GetConfigs()
	api := "/v1/auth/token/"
	path := config.GatewayHost + api
	AuthToken = ""
	var request = map[string]string{}
	request["app_id"] = config.MiAppID
	request["app_key"] = config.MiAppKey
	request["app_secret"] = config.MiAppSecret
	response := utils.HTTPRequest(path, "POST", request, "")
	if response.Code != 200 {
		fmt.Println(api + "：授权出错！")
		return
	}
	AuthToken = response.Data.(string)
}
