package services

import (
	"fmt"
	"kefu_go_robot/conf"
	"kefu_server/utils"
)

// MessageRepository struct
type MessageRepository struct{}

// GetMessageRepositoryInstance get instance
func GetMessageRepositoryInstance() *MessageRepository {
	instance := new(MessageRepository)
	return instance
}

// PushMessage Push Message
func (r *MessageRepository) PushMessage(payload string) bool {
	config := new(conf.Cionfigs).GetConfigs()
	api := "/v1/public/message/push"
	path := config.GatewayHost + api
	var request = map[string]string{}
	request["msgType"] = "NORMAL_MSG"
	request["payload"] = payload
	response := utils.HTTPRequest(path, "POST", request, AuthToken)
	if response.Code != 200 {
		fmt.Println(response)
		return false
	}
	return true
}
