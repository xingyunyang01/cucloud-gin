package cgin

import (
	"fmt"
	"reflect"
	"strings"
)

// 注解处理
type Annotation interface {
	SetTag(tag reflect.StructTag)
}

var AnnotationList []Annotation

// 判断当前进入对象是否是注解
func IsAnnotation(t reflect.Type) bool {
	fmt.Println(len(AnnotationList))
	for _, annotation := range AnnotationList { //遍历Annotation列表
		if reflect.TypeOf(annotation) == t { //是否传入的类型在Annotation列表列表里
			return true
		}
	}

	return false
}

func init() {
	AnnotationList = make([]Annotation, 0)
	AnnotationList = append(AnnotationList, new(Value))
}

type Value struct {
	tag         reflect.StructTag
	Beanfactory *BeanFactory
}

func (this *Value) SetTag(tag reflect.StructTag) {
	this.tag = tag
}

func (this *Value) String() string {
	get_prefix := this.tag.Get("prefix") //获取tag值
	if get_prefix == "" {
		return ""
	}
	prefix := strings.Split(get_prefix, ".") //把tag值按.进行分割，放入切片中
	if config := this.Beanfactory.GetBean(new(SysConfig)); config != nil {
		get_value := GetConfigValue(config.(*SysConfig).Config, prefix, 0)
		if get_value != nil {
			return fmt.Sprintf("%v", get_value)
		} else {
			return ""
		}
	} else {
		return ""
	}
}
