package robot

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"kefu_go_robot/conf"
	"net/http"
	"strconv"
	"strings"

	"github.com/astaxie/beego/logs"
)

// NewTokenHandler ...
func NewTokenHandler(appAccount string) *TokenHandler {
	config := new(conf.Cionfigs).GetConfigs()
	tokenHandler := new(TokenHandler)
	tokenHandler.httpURL = config.MiHost
	tokenHandler.AppID, _ = strconv.ParseInt(config.MiAppID, 10, 64)
	tokenHandler.AppKey = config.MiAppKey
	tokenHandler.AppSecret = config.MiAppSecret
	tokenHandler.AppAccount = appAccount
	return tokenHandler
}

// FetchToken ...
func (c *TokenHandler) FetchToken() *string {
	jsonBytes, err := json.Marshal(*c)
	if err != nil {
		logs.Info("FetchToken error==%v", err)
		return nil
	}
	requestJSONBodygo := bytes.NewBuffer(jsonBytes).String()
	request, err := http.Post(c.httpURL, "application/json", strings.NewReader(requestJSONBodygo))
	if err != nil {
		logs.Info("http get FetchToken error==%v", err)
		return nil
	}
	defer request.Body.Close()
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		logs.Info("ioutil.ReadAll(request.Body) FetchToken error==%v", err)
		return nil
	}
	token := string(body)
	return &token
}
