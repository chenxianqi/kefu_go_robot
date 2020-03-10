package conf

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type conf struct {
	Host      string `yaml:"Host"`
	AppID     string `yaml:"AppID"`
	AppKey    string `yaml:"AppKey"`
	AppSecret string `yaml:"AppSecret"`
}

func (c *conf) getConf() *conf {
	yamlFile, err := ioutil.ReadFile("conf.yaml")
	if err != nil {
		fmt.Println(err.Error())
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		fmt.Println(err.Error())
	}
	return c
}
