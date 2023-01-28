package cgin

import (
	"log"

	"github.com/xingyunyang01/cucloud-gin/cgin/utils"
	"gopkg.in/yaml.v3"
)

type ServerConfig struct {
	Port int32
	Name string
	Html string
}

type UserConfig map[string]interface{}

// 递归读取用户配置文件
func GetConfigValue(m UserConfig, prefix []string, index int) interface{} {
	key := prefix[index]
	if v, ok := m[key]; ok {
		if index == len(prefix)-1 { //到了最后一个
			return v
		} else {
			index = index + 1
			if mv, ok := v.(UserConfig); ok { //值必须是UserConfig类型
				return GetConfigValue(mv, prefix, index)
			} else {
				return nil
			}

		}
	}
	return nil
}

// 系统配置
type SysConfig struct {
	Server *ServerConfig
	Config UserConfig
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
