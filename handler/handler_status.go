package handler

import "fmt"

// StatusHandler struct
type StatusHandler struct {
	appAccount string
}

// NewStatusHandler newStatusHandler
func NewStatusHandler(appAccount string) *StatusHandler {
	return &StatusHandler{appAccount}
}

// HandleChange handleChange
func (c StatusHandler) HandleChange(isOnline bool, errType, errReason, errDescription *string) {
	if isOnline {
		fmt.Println("机器人霸道上线 status changed: online.")
	} else {
		// 有机器人掉线，重新登录
		fmt.Printf("[机器人挂掉了] status changed: offline，errType:%v, errReason:%v, errDes:%v\r\n", *errType, *errReason, *errDescription)
	}
}
