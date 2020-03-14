package services

import (
	"context"
	"fmt"
	"kefu_go_robot/grpcc"
	"kefu_server/grpcs"
	"kefu_server/models"
	"kefu_server/utils"
)

// AdminRepository struct
type AdminRepository struct{}

// GetAdminRepositoryInstance get instance
func GetAdminRepositoryInstance() *AdminRepository {
	instance := new(AdminRepository)
	return instance
}

// GetOnlineAdmins get online all admin
func (r *AdminRepository) GetOnlineAdmins() []models.Admin {
	grpcClient := grpcc.GrpcClient()
	res, err := grpcClient.GetOnlineAdmins(context.Background(), &grpcs.Request{Data: ""})
	if err != nil {
		fmt.Printf("SearchKnowledgeTitles get titles res==%v", err)
		return nil
	}
	var admins []models.Admin
	utils.StringToInterface(res.Data, &admins)
	return admins
}
