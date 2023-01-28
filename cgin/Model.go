package cgin

import (
	"encoding/json"
	"log"
)

type Model interface {
	String() string
}

// 支持多Model的形式，之所以把Models的类型定义为string，而不是[]Model，主要是为了简单
type Models string

func MakeModels(v interface{}) Models { //将[]Model序列化成[]byte，之后转成string
	b, err := json.Marshal(v)
	if err != nil {
		log.Println(err)
	}

	return Models(b)
}
