package services

import (
	"context"
	"fmt"
	"kefu_go_robot/grpcc"
	"kefu_server/grpcs"
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
		fmt.Printf("PushNewContacts Contact err==%v", err)
	}
}
