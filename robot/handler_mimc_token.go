package robot

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"kefu_go_robot/conf"
	"net/http"
	"strconv"
	"strings"
)

// TokenHandler ...
type TokenHandler struct {
	httpURL    string
	AppID      int64  `json:"appId"`
	AppKey     string `json:"appKey"`
	AppSecret  string `json:"appSecret"`
	AppAccount string `json:"appAccount"`
}

// GetMiMcToken ...
func GetMiMcToken(accountID string) (string, error) {
	config := new(conf.Cionfigs).GetConfigs()
	tokenHandler := new(TokenHandler)
	tokenHandler.httpURL = config.MiHost
	tokenHandler.AppID, _ = strconv.ParseInt(config.MiAppID, 10, 64)
	tokenHandler.AppKey = config.MiAppKey
	tokenHandler.AppSecret = config.MiAppSecret
	tokenHandler.AppAccount = accountID
	jsonBytes, err := json.Marshal(*tokenHandler)
	if err != nil {
		return "", err
	}
	requestJSONBody := bytes.NewBuffer(jsonBytes).String()
	request, err := http.Post(tokenHandler.httpURL, "application/json", strings.NewReader(requestJSONBody))
	if err != nil {
		return "", err
	}
	defer request.Body.Close()
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		return "", err
	}
	token := string(body)
	return token, nil
}
