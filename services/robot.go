package services

import (
	"context"
	"kefu_go_robot/grpcc"
	"kefu_server/grpcs"
	"kefu_server/models"
	"kefu_server/utils"
	"log"
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
	grpcClient, err := grpcc.GrpcClient()
	if err != nil {
		log.Fatalf("GetOnlineAllRobots auth use grpcClient err==%v", err)
	}
	res, _ := grpcClient.GetOnlineAllRobots(context.Background(), &grpcs.Request{Data: ""})
	var robots []*models.Robot
	utils.StringToInterface(res.Data, &robots)
	return robots
}
