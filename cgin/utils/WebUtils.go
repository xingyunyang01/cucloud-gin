package utils

import (
	"io/ioutil"
	"log"
	"os"
	"runtime"
)

func LoadConfigFile() []byte {
	sysType := runtime.GOOS
	dir, _ := os.Getwd() //获取程序运行的根目录
	var file string

	if sysType == "linux" {
		file = dir + "/sys.yaml"
	} else if sysType == "windows" {
		file = dir + "\\sys.yaml"
	}
	b, err := ioutil.ReadFile(file)
	if err != nil {
		log.Println(err)
		return nil
	}

	return b
}
