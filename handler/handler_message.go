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

	msg "github.com/Xiaomi-mimc/mimc-go-sdk/message"
)

// HandleMessage ...
func (c MsgHandler) HandleMessage(packets *list.List) {
	for ele := packets.Front(); ele != nil; ele = ele.Next() {

		// 收到的原始消息
		p2pMsg := ele.Value.(*msg.P2PMessage)

		// get message
		var message models.Message
		msgContentByte, err := base64.StdEncoding.DecodeString(string(p2pMsg.Payload()))
		err = json.Unmarshal(msgContentByte, &message)

		fmt.Printf("get message type %v", message.BizType)

		// 当前服务机器人
		// var mcUserRobot *mimc.MCUser
		robot := GetRunRobotInfo(message.ToAccount)
		if robot != nil && robot.ID == message.FromAccount {
			return
		}

		// message.BizType
		MessageRepository := services.GetMessageRepositoryInstance()
		switch message.BizType {

		// 消息入库
		case "into":
			messageString := utils.InterfaceToString(message.Payload)
			MessageRepository.PushMessage(messageString)
			return

		// 撤销消息
		case "cancel":
			key, _ := strconv.ParseInt(message.Payload, 10, 64)
			MessageRepository.CancelMessage(
				models.RemoveMessageRequestDto{
					FromAccount: message.FromAccount,
					ToAccount:   message.ToAccount,
					Key:         key,
				})
			return
		// 搜索知识库
		case "search_knowledge":
			return
		// 与机器人握手
		case "handshake":
			return
		default:

		}

		if err == nil {
			// MessageP2P(message)
		}

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
