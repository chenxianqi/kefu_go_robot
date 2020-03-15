package conf

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// Cionfigs struct
type Cionfigs struct {
	MiHost      string `yaml:"mi_host"`
	MiAppID     string `yaml:"mi_app_id"`
	MiAppKey    string `yaml:"mi_app_key"`
	MiAppSecret string `yaml:"mi_app_secret"`
	GatewayHost string `yaml:"gateway_host"`
	GrpcHost    string `yaml:"grpc_host"`
	GrpcPort    string `yaml:"grpc_port"`
}

// GetConfigs instance
func (c *Cionfigs) GetConfigs() *Cionfigs {
	yamlFile, err := ioutil.ReadFile("conf/conf.yaml")
	if err != nil {
		fmt.Println(err.Error())
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		fmt.Println(err.Error())
	}
	return c
}
