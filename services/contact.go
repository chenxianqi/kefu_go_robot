package services

import (
	"context"
	"kefu_go_robot/grpcc"
	"kefu_server/grpcs"

	"github.com/astaxie/beego/logs"
)

// ContactRepository struct
type ContactRepository struct{}

// GetContactRepositoryInstance get instance
func GetContactRepositoryInstance() *ContactRepository {
	instance := new(ContactRepository)
	return instance
}

// PushNewContacts Contact
func (r *ContactRepository) PushNewContacts(uid string) {
	grpcClient := grpcc.GrpcClient()
	_, err := grpcClient.PushNewContacts(context.Background(), &grpcs.Request{Data: uid})
	if err != nil {
		logs.Info("PushNewContacts Contact err==%v", err)
	}
}
