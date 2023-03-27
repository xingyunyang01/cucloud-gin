package cgin

import (
	"log"

	"github.com/xingyunyang01/cucloud-gin/cgin/utils"
	"gopkg.in/yaml.v3"
)

type ServerConfig struct {
	Port int32
	Name string
}

// 系统配置
type SysConfig struct {
	Server *ServerConfig
}

func NewSysConfig() *SysConfig {
	return &SysConfig{Server: &ServerConfig{Port: 8080, Name: "cucloud-gin"}}
}

func InitConfig() *SysConfig {
	config := NewSysConfig()

	if b := utils.LoadConfigFile(); b != nil { //读取application.yaml文件内容
		err := yaml.Unmarshal(b, config) //将byte反序列化成结构体
		if err != nil {
			log.Fatal(err)
		}
	}
	return config
}
