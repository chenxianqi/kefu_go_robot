package handler

import (
	"kefu_go_robot/conf"
	"kefu_go_robot/services"
	"kefu_server/models"
	"strconv"

	"github.com/Xiaomi-mimc/mimc-go-sdk"
)

// NewMsgHandler ...
func NewMsgHandler(appAccount string) *MsgHandler {
	return &MsgHandler{appAccount}
}

// MsgHandler ...
type MsgHandler struct {
	appAccount string
}

// MCUserRobots 工作中的机器人
var MCUserRobots []*mimc.MCUser

// Robots 机器人资料列表
var Robots []*models.Robot

// CreateRobot 创建机器人
func CreateRobot(appAccount string) *mimc.MCUser {
	config := new(conf.Cionfigs).GetConfigs()
	appID, _ := strconv.ParseInt(config.MiAppID, 10, 64)
	mcUser := mimc.NewUser(uint64(appID), appAccount)
	mcUser.RegisterStatusDelegate(NewStatusHandler(appAccount))
	mcUser.RegisterTokenDelegate(NewTokenHandler(appAccount))
	mcUser.RegisterMessageDelegate(NewMsgHandler(appAccount))
	mcUser.InitAndSetup()
	return mcUser
}

// GetRunRobotInfo Get current robot
func GetRunRobotInfo(id int64) *models.Robot {
	for _, robot := range Robots {
		if robot.ID == id {
			return robot
		}
	}
	return nil
}

// GetOnlineRobots get robot all
func GetOnlineRobots() []*models.Robot {
	RobotRepository := services.GetRobotRepositoryInstance()
	robots := RobotRepository.GetOnlineAllRobots()
	return robots
}

// RobotRun init
func RobotRun() {

	// Log out if any robot is working
	if len(MCUserRobots) > 0 {
		for _, robot := range MCUserRobots {
			robot.Logout()
			robot.Destory()
		}
		MCUserRobots = []*mimc.MCUser{}
	}
	Robots = GetOnlineRobots()
	var tempRobots []*mimc.MCUser
	for _, robot := range Robots {
		if robot.Switch == 1 {
			rb := CreateRobot(strconv.FormatInt(robot.ID, 10))
			tempRobots = append(tempRobots, rb)
			rb.Login()
		}
	}
	MCUserRobots = tempRobots

}
