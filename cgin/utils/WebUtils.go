package utils

import (
	"io/ioutil"
	"log"
	"os"
)

func LoadConfigFile() []byte {
	dir, _ := os.Getwd() //获取程序运行的根目录
	file := dir + "\\application.yaml"
	b, err := ioutil.ReadFile(file)
	if err != nil {
		log.Println(err)
		return nil
	}
	return b
}
