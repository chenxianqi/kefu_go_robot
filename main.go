package main

import (
	"kefu_go_robot/grpcc"
	"kefu_go_robot/robot"
	"time"

	"github.com/astaxie/beego/logs"
)

func main() {

	// robot log
	logs.SetLogger(logs.AdapterFile, `{"filename":"project_robot.log","level":6,"maxlines":0,"maxsize":0,"daily":true,"maxdays":10,"color":true}`)
	logs.EnableFuncCallDepth(true)

	// grpcc init
	grpcc.Run()

	// RobotRun
	robot.Run()

	// Restart all robots every 60 minutes
	c := time.Tick(60 * 60 * time.Second)
	for {
		<-c
		go robot.Run()
	}

}
