package services

import (
	"context"
	"kefu_go_robot/grpcc"
	"kefu_server/grpcs"
	"kefu_server/models"
	"kefu_server/utils"
	"log"
)

// MessageRepository struct
type MessageRepository struct{}

// GetMessageRepositoryInstance get instance
func GetMessageRepositoryInstance() *MessageRepository {
	instance := new(MessageRepository)
	return instance
}

// PushMessage Push Message
func (r *MessageRepository) PushMessage(payload string) {
	grpcClient, err := grpcc.GrpcClient()
	if err != nil {
		log.Fatalf("PushMessage Push Message use grpcClient err==%v", err)
	}
	grpcClient.PushMessage(context.Background(), &grpcs.Request{Data: payload})
}

// CancelMessage Cancel Message
func (r *MessageRepository) CancelMessage(request models.RemoveMessageRequestDto) {
	grpcClient, err := grpcc.GrpcClient()
	if err != nil {
		log.Fatalf("PushMessage Push Message use grpcClient err==%v", err)
	}
	grpcClient.CancelMessage(context.Background(), &grpcs.Request{Data: utils.InterfaceToString(request)})
}
