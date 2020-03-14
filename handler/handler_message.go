package handler

import (
	"container/list"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"kefu_go_robot/services"
	"kefu_server/models"
	"kefu_server/utils"
	"strconv"
	"strings"
	"time"

	"github.com/Xiaomi-mimc/mimc-go-sdk"
	msg "github.com/Xiaomi-mimc/mimc-go-sdk/message"
)

// create a message
func createMessage(
	bizType string,
	fromAccount int64,
	toAccount int64,
	payload string,
	transferAccount int64,
	platform int64,
	read int,
) models.Message {
	return models.Message{
		FromAccount:     fromAccount,
		ToAccount:       toAccount,
		BizType:         bizType,
		Timestamp:       time.Now().Unix() + 1,
		Key:             time.Now().Unix(),
		TransferAccount: transferAccount,
		Platform:        platform,
		Payload:         payload,
		Read:            read,
	}
}

// adminData struct
type adminData struct {
	ID       int64  `json:"id"`
	NickName string `json:"nickname"`
	Avatar   string `json:"avatar"`
}

// HandleMessage ...
func (c MsgHandler) HandleMessage(packets *list.List) {
	for ele := packets.Front(); ele != nil; ele = ele.Next() {

		// 收到的原始消息
		p2pMsg := ele.Value.(*msg.P2PMessage)

		// get message
		var message models.Message
		msgContentByte, _ := base64.StdEncoding.DecodeString(string(p2pMsg.Payload()))
		json.Unmarshal(msgContentByte, &message)

		// 消息入库
		if message.BizType == "into" {
			messageString := utils.InterfaceToString(message.Payload)
			services.GetMessageRepositoryInstance().PushMessage(messageString)
			return
		}

		fmt.Printf("get message %v \r\n", message)

		// 当前服务机器人
		var mcUserRobot *mimc.MCUser
		robot := GetRunRobotInfo(message.ToAccount)
		if robot == nil {
			fmt.Print("robot info == nil \r\n")
			return
		}
		if robot != nil && robot.ID == message.FromAccount {
			return
		}
		for _, mcRbt := range MCUserRobots {
			rbtID, _ := strconv.ParseInt(mcRbt.AppAccount(), 10, 64)
			if robot != nil && rbtID == robot.ID {
				mcUserRobot = mcRbt
				break
			}
		}

		// 撤销消息
		if message.BizType == "cancel" {
			key, _ := strconv.ParseInt(message.Payload, 10, 64)
			services.GetMessageRepositoryInstance().CancelMessage(
				models.RemoveMessageRequestDto{
					FromAccount: message.FromAccount,
					ToAccount:   message.ToAccount,
					Key:         key,
				})
			return
		}
		// 返回给对方的消息内容
		var messageContent string = ""
		var bizType string = "text"
		knowledgeBaseRepository := services.GetKnowledgeBaseRepositoryInstance()
		// 搜索知识库
		if message.BizType == "search_knowledge" {
			bizType = "search_knowledge"
			payload := strings.Trim(strings.ToLower(message.Payload), " ")
			if payload == "" {
				fmt.Print("payload is empty")
				return
			}
			knowledgeTitles := knowledgeBaseRepository.SearchKnowledgeTitles(models.KnowledgeBaseTitleRequestDto{
				Payload:     payload,
				KeyWords:    robot.KeyWord,
				Platform:    message.Platform,
				IsSerachSub: false,
				Limit:       5,
			})
			knowledgeTitlesByte, _ := json.Marshal(knowledgeTitles)
			messageContent = string(knowledgeTitlesByte)

		} else if message.BizType == "handshake" {

			// 与机器人握手
			messageContent = robot.Welcome
			bizType = "welcome"

		} else {

			// 判断是否符合转人工
			artificial := strings.Split(strings.Trim(robot.Artificial, "|"), "|")
			isTransfer := false
			if message.Payload == "人工" {
				isTransfer = true
			} else {
				for i := 0; i < len(artificial); i++ {
					if artificial[i] == message.Payload {
						isTransfer = true
						break
					}
				}
			}
			// 符合
			if isTransfer {
				admins := services.GetAdminRepositoryInstance().GetOnlineAdmins()
				if len(admins) <= 0 {
					messageContent = robot.NoServices
				} else {
					// 平均分配客服
					admin := admins[0]
					adminDataJSON, _ := json.Marshal(adminData{ID: admin.ID, NickName: admin.NickName, Avatar: admin.Avatar})
					messageContent = string(adminDataJSON)

					// 发送一条消息告诉客服端
					var newMsgBase64 string
					newMsg := models.Message{}
					newMsg.BizType = "transfer"
					newMsg.FromAccount = message.FromAccount
					newMsg.ToAccount = admin.ID
					newMsg.Timestamp = time.Now().Unix()
					newMsg.TransferAccount = admin.ID
					newMsg.Payload = "系统将客户分配给您"
					newMsgBase64 = utils.InterfaceToString(newMsg)

					// 发送与消息入库
					services.GetMessageRepositoryInstance().PushMessage(utils.InterfaceToString(newMsgBase64))
					mcUserRobot.SendMessage(strconv.FormatInt(admin.ID, 10), []byte(newMsgBase64))

					newMsg.FromAccount = robot.ID
					newMsg.ToAccount = message.FromAccount
					newMsg.Payload = messageContent
					newMsgBase64 = utils.InterfaceToString(newMsg)

					// 消息入库
					mcUserRobot.SendMessage(strconv.FormatInt(message.FromAccount, 10), []byte(newMsgBase64))

					// 帮助客服发送欢迎语
					newMsg.BizType = "text"
					newMsg.Payload = admin.AutoReply
					newMsg.ToAccount = message.FromAccount
					newMsg.FromAccount = admin.ID
					newMsgBase64 = utils.InterfaceToString(newMsg)

					// 发送与消息入库
					mcUserRobot.SendMessage(strconv.FormatInt(admin.ID, 10), []byte(newMsgBase64))
					mcUserRobot.SendMessage(strconv.FormatInt(message.FromAccount, 10), []byte(newMsgBase64))
					services.GetMessageRepositoryInstance().PushMessage(utils.InterfaceToString(newMsgBase64))

					// 推送列表给客服
					services.GetContactRepositoryInstance().PushNewContacts(strconv.FormatInt(admin.ID, 10))

					// 转接入库用于统计服务次数
					servicesStatistical := models.ServicesStatistical{UserAccount: message.FromAccount, ServiceAccount: admin.ID, Platform: message.Platform, TransferAccount: robot.ID, CreateAt: time.Now().Unix()}
					services.GetStatisticalRepositoryInstance().Add(servicesStatistical)
					return
				}

			} else {

				// 完全匹配知识库
				_knowledgeBase := knowledgeBaseRepository.GetKnowledgeBaseWithTitleAndPlatform(message.Payload, message.Platform)
				if _knowledgeBase != nil {
					bizType = "text"
					messageContent = _knowledgeBase.Content
				} else {

					bizType = "knowledge"
					// 找主标题
					knowledgeTitles := knowledgeBaseRepository.SearchKnowledgeTitles(models.KnowledgeBaseTitleRequestDto{
						Payload:     strings.Trim(strings.ToLower(message.Payload), " "),
						KeyWords:    "",
						IsSerachSub: false,
						Platform:    message.Platform,
						Limit:       4,
					})
					if len(knowledgeTitles) > 0 {
						messageContentByte, _ := json.Marshal(knowledgeTitles)
						messageContent = string(messageContentByte)
					} else {

						// 找副标题
						knowledgeTitles := knowledgeBaseRepository.SearchKnowledgeTitles(models.KnowledgeBaseTitleRequestDto{
							Payload:     strings.Trim(strings.ToLower(message.Payload), " "),
							KeyWords:    "",
							IsSerachSub: true,
							Platform:    message.Platform,
							Limit:       4,
						})
						if len(knowledgeTitles) > 0 {
							messageContentByte, _ := json.Marshal(knowledgeTitles)
							messageContent = string(messageContentByte)
						} else {
							knowledgeTitles := knowledgeBaseRepository.SearchKnowledgeTitles(models.KnowledgeBaseTitleRequestDto{
								Payload:     strings.Trim(strings.ToLower(message.Payload), " "),
								KeyWords:    robot.KeyWord,
								Platform:    message.Platform,
								IsSerachSub: true,
								Limit:       4,
							})
							if len(knowledgeTitles) > 0 {
								bizType = "knowledge"
								messageContentByte, _ := json.Marshal(knowledgeTitles)
								messageContent = string(messageContentByte)
							} else {
								bizType = "text"
								messageContent = robot.Understand
							}
						}

					}

				}

			}

		}
		// 消息体
		_message := models.Message{}
		_message.BizType = bizType
		_message.FromAccount = message.ToAccount
		_message.Timestamp = time.Now().Unix() + 1
		_message.ToAccount = message.FromAccount
		_message.Key = time.Now().Unix()
		_message.Platform = message.Platform
		_message.Payload = messageContent
		messageString := utils.InterfaceToString(_message)

		// 发给用户
		mcUserRobot.SendMessage(strconv.FormatInt(message.FromAccount, 10), []byte(messageString))

		// 消息入库
		_messageString := utils.InterfaceToString(messageString)
		services.GetMessageRepositoryInstance().PushMessage(_messageString)

	}

}

// HandleGroupMessage ...
func (c MsgHandler) HandleGroupMessage(packets *list.List) {
	//for ele := packets.Front(); ele != nil; ele = ele.Next() {
	//	p2tmsg := ele.Value.(*msg.P2TMessage)
	//	logger.Info("[%v] [handle p2t msg]%v  -> %v: %v, pcktId: %v, timestamp: %v.", c.appAccount, *(p2tmsg.FromAccount()), *(p2tmsg.GroupId()), string(p2tmsg.Payload()), *(p2tmsg.PacketId()), *(p2tmsg.Timestamp()))
	//}
}

// HandleServerAck ...
func (c MsgHandler) HandleServerAck(packetID *string, sequence, timestamp *int64, errMsg *string) {
	//logs.Info("[%v] [handle server ack] packetId:%v, seqId: %v, timestamp:%v.", c.appAccount, *packetId, *sequence, *timestamp)
}

// HandleSendMessageTimeout ...
func (c MsgHandler) HandleSendMessageTimeout(message *msg.P2PMessage) {
	//logs.Info("[%v] [handle p2pmsg timeout] packetId:%v, msg:%v, time: %v.", c.appAccount, *(message.PacketId()), string(message.Payload()), time.Now())
}

// HandleSendGroupMessageTimeout ...
func (c MsgHandler) HandleSendGroupMessageTimeout(message *msg.P2TMessage) {
	// logger.Info("[%v] [handle p2tmsg timeout] packetId:%v, msg:%v.", c.appAccount, *(message.PacketId()), string(message.Payload()))
}
