package main

import (
	"kefu_go_robot/grpcc"
	"kefu_go_robot/handler"
)

func main() {

	// Use semaphores to cooperate with channels to achieve permanent blocking and keep alive processes
	end := make(chan bool, 1)

	// // FetchToken
	// AuthTokenRepository := services.GetAuthTokenRepositoryInstance()
	// AuthTokenRepository.FetchToken()
	// if services.AuthToken == "" {
	// 	fmt.Println("/v1/auth/token/：授权出错！")
	// 	return
	// }
	// // AuthToken
	// fmt.Println("AuthToken==", services.AuthToken)

	// grpcc init
	grpcc.Run()

	// RobotRun
	handler.RobotRun()

	// The blocking will stop when we pass the SIGUSR2 signal (kill -USR2 [pid])

	<-end
}
