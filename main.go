package main

import (
	"kefu_go_robot/grpcc"
	"kefu_go_robot/robot"
	"time"

	"github.com/Xiaomi-mimc/mimc-go-sdk/util/log"
)

func main() {

	// robot log
	log.SetLogLevel(log.FatalLevel)
	// log.SetLogLevel(log.ErrorLevel)

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
