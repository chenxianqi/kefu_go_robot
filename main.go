package main

import (
	"fmt"
	"kefu_go_robot/auth"
	"kefu_go_robot/conf"
)

func main() {

	auth.FetchToken()
	fmt.Println(auth.Token)

	config := new(conf.Cionfigs).GetConfigs()
	fmt.Println(config.MiAppID)
	fmt.Println(config.MiAppKey)
	fmt.Println(config.MiAppSecret)
	fmt.Println(config.MiHost)

}
