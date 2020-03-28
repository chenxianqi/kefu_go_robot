package services

import (
	"context"
	"kefu_go_robot/grpcc"
	"kefu_server/grpcs"
	"kefu_server/models"
	"kefu_server/utils"

	"github.com/astaxie/beego/logs"
)

// UserRepository struct
type UserRepository struct{}

// GetUserRepositoryInstance get instance
func GetUserRepositoryInstance() *UserRepository {
	instance := new(UserRepository)
	return instance
}

// Update update user
func (r *UserRepository) Update(user models.User) {
	grpcClient := grpcc.GrpcClient()
	_, err := grpcClient.UpdateUser(context.Background(), &grpcs.Request{Data: utils.InterfaceToString(user)})
	if err != nil {
		logs.Info("Update update user==%v", err)
	}
}
