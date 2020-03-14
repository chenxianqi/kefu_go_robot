package services

import (
	"context"
	"fmt"
	"kefu_go_robot/grpcc"
	"kefu_server/grpcs"
	"kefu_server/models"
	"kefu_server/utils"
)

// StatisticalRepository struct
type StatisticalRepository struct{}

// GetStatisticalRepositoryInstance get instance
func GetStatisticalRepositoryInstance() *StatisticalRepository {
	instance := new(StatisticalRepository)
	return instance
}

// Add Statistical
func (r *StatisticalRepository) Add(request models.ServicesStatistical) {
	grpcClient := grpcc.GrpcClient()
	_, err := grpcClient.PushMessage(context.Background(), &grpcs.Request{Data: utils.InterfaceToString(request)})
	if err != nil {
		fmt.Printf("Add Statistical err==%v", err)
	}
}
