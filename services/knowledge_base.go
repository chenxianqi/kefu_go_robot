package services

import (
	"context"
	"encoding/json"
	"fmt"
	"kefu_go_robot/grpcc"
	"kefu_server/grpcs"
	"kefu_server/models"
	"kefu_server/utils"
	"strconv"
)

// KnowledgeBaseRepository struct
type KnowledgeBaseRepository struct{}

// GetKnowledgeBaseRepositoryInstance get instance
func GetKnowledgeBaseRepositoryInstance() *KnowledgeBaseRepository {
	instance := new(KnowledgeBaseRepository)
	return instance
}

// GetKnowledgeBaseWithTitleAndPlatform get with title
func (r *KnowledgeBaseRepository) GetKnowledgeBaseWithTitleAndPlatform(title string, platform int64) *models.KnowledgeBase {
	request := make(map[string]string)
	request["title"] = title
	request["platform"] = strconv.FormatInt(platform, 10)
	byteData, _ := json.Marshal(request)
	grpcClient := grpcc.GrpcClient()
	res, err := grpcClient.GetKnowledgeBaseWithTitleAndPlatform(context.Background(), &grpcs.Request{Data: string(byteData)})
	if err != nil {
		fmt.Printf("SearchKnowledgeTitles get titles res==%v", err)
	}
	var knowledgeBase *models.KnowledgeBase
	utils.StringToInterface(res.Data, &knowledgeBase)
	return knowledgeBase
}

// SearchKnowledgeTitles get titles
func (r *KnowledgeBaseRepository) SearchKnowledgeTitles(request models.KnowledgeBaseTitleRequestDto) []models.KnowledgeBaseTitleDto {
	grpcClient := grpcc.GrpcClient()
	res, err := grpcClient.SearchKnowledgeTitles(context.Background(), &grpcs.Request{Data: utils.InterfaceToString(request)})
	if err != nil {
		fmt.Printf("SearchKnowledgeTitles get titles res==%v", err)
		return nil
	}
	var KnowledgeBaseTitles []models.KnowledgeBaseTitleDto
	utils.StringToInterface(res.Data, &KnowledgeBaseTitles)
	if KnowledgeBaseTitles == nil {
		return []models.KnowledgeBaseTitleDto{}
	}
	return KnowledgeBaseTitles
}
