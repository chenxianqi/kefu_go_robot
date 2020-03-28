package services

import (
	"context"
	"kefu_go_robot/grpcc"
	"kefu_server/grpcs"
	"kefu_server/models"
	"kefu_server/utils"

	"github.com/astaxie/beego/logs"
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
	grpcClient := grpcc.GrpcClient()
	res, err := grpcClient.GetOnlineAllRobots(context.Background(), &grpcs.Request{Data: ""})
	if err != nil {
		logs.Info("GetOnlineAllRobots auth use grpcClient res==%v", err)
		return nil
	}
	var robots []*models.Robot
	utils.StringToInterface(res.Data, &robots)
	return robots
}
