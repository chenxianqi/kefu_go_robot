package robot

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
		fmt.Printf("机器人:%v 霸道上线 status changed: online.", c.appAccount)
	} else {
		// 有机器人掉线，重新登录
		fmt.Printf("[机器人:%v 挂掉了] status changed: offline，errType:%v, errReason:%v, errDes:%v\r\n", c.appAccount, *errType, *errReason, *errDescription)
	}
}
