package services

import (
	"encoding/json"
	"fmt"
	"kefu_go_robot/conf"
	"kefu_server/models"
	"kefu_server/utils"
)

// RobotRepository struct
type RobotRepository struct{}

// GetRobotRepositoryInstance get instance
func GetRobotRepositoryInstance() *RobotRepository {
	instance := new(RobotRepository)
	return instance
}

// GetOnlineAllRobots auth
func (r *RobotRepository) GetOnlineAllRobots() []*models.Robot {
	config := new(conf.Cionfigs).GetConfigs()
	api := "/v1/robot/online/all"
	path := config.GatewayHost + api
	response := utils.HTTPRequest(path, "GET", "", AuthToken)
	if response.Code != 200 {
		fmt.Println(api + "：授权出错！")
		return nil
	}
	var robots []*models.Robot
	listData := response.Data.([]interface{})
	mapByte, _ := json.Marshal(listData)
	_ = json.Unmarshal(mapByte, &robots)
	return robots
}
