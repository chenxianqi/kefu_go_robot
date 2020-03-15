package services

import (
	"context"
	"fmt"
	"kefu_go_robot/grpcc"
	"kefu_server/grpcs"
	"kefu_server/models"
	"kefu_server/utils"
)

// MessageRepository struct
type MessageRepository struct{}

// GetMessageRepositoryInstance get instance
func GetMessageRepositoryInstance() *MessageRepository {
	instance := new(MessageRepository)
	return instance
}

// InsertMessage Push Message
func (r *MessageRepository) InsertMessage(payload string) {
	grpcClient := grpcc.GrpcClient()
	_, err := grpcClient.InsertMessage(context.Background(), &grpcs.Request{Data: payload})
	if err != nil {
		fmt.Printf("InsertMessage Push Message err==%v", err)
	}
}

// CancelMessage Cancel Message
func (r *MessageRepository) CancelMessage(request models.RemoveMessageRequestDto) {
	grpcClient := grpcc.GrpcClient()
	_, err := grpcClient.CancelMessage(context.Background(), &grpcs.Request{Data: utils.InterfaceToString(request)})
	if err != nil {
		fmt.Printf("CancelMessage Cancel Message err==%v", err)
	}
}
