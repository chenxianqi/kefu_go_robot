package main

import (
	"fmt"
	"kefu_go_robot/grpcc"
	"kefu_go_robot/handler"
	"kefu_go_robot/services"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	// Use semaphores to cooperate with channels to achieve permanent blocking and keep alive processes
	sig := make(chan os.Signal, 1)

	// FetchToken
	AuthTokenRepository := services.GetAuthTokenRepositoryInstance()
	AuthTokenRepository.FetchToken()
	if services.AuthToken == "" {
		fmt.Println("/v1/auth/token/：授权出错！")
		return
	}

	// AuthToken
	fmt.Println("AuthToken==", services.AuthToken)

	// grpcc init
	grpcc.Run()

	// RobotRun
	handler.RobotRun()

	// The blocking will stop when we pass the SIGUSR2 signal (kill -USR2 [pid])
	signal.Notify(sig, syscall.SIGUSR2)
	<-sig
}
