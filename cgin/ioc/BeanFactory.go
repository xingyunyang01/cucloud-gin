package ioc

import (
	"reflect"

	"github.com/shenyisyn/goft-expr/src/expr"
)

var BeanFactory *BeanFactoryImpl

func init() {
	BeanFactory = NewBeanFactory()
}

type BeanFactoryImpl struct {
	beanMapper BeanMapper
	ExprMap    map[string]interface{}
}

func NewBeanFactory() *BeanFactoryImpl {
	return &BeanFactoryImpl{beanMapper: make(BeanMapper), ExprMap: make(map[string]interface{})}
}

func (this *BeanFactoryImpl) GetBeanMapper() BeanMapper {
	return this.beanMapper
}

// 向beanMapper添加内容
func (this *BeanFactoryImpl) Set(vList ...interface{}) {
	if vList == nil || len(vList) == 0 {
		return
	}

	for _, v := range vList {
		this.beanMapper.add(v)
	}
}

// 从beanMapper获取内容
func (this *BeanFactoryImpl) Get(bean interface{}) interface{} {
	if bean == nil {
		return nil
	}

	v := this.beanMapper.get(bean)
	if v.IsValid() {
		return v.Interface()
	}

	return nil
}

// 注入依赖
// 参数bean需要传入一个指针，指针指向的内容是结构体
func (this *BeanFactoryImpl) Apply(bean interface{}) {
	if bean == nil {
		return //如果bean是空则直接返回
	}

	//获取bean的值
	v := reflect.ValueOf(bean)
	if v.Kind() == reflect.Ptr {
		v = v.Elem() //获取指针指向的内容
	}
	if v.Kind() != reflect.Struct {
		return //如果指针指向的内容不是结构体则返回
	}

	//遍历结构体成员
	for i := 0; i < v.NumField(); i++ {
		field := v.Type().Field(i) //得到成员

		if v.Field(i).CanSet() && field.Tag.Get("inject") != "" {
			if field.Tag.Get("inject") != "-" { //多例模式(表达式的方式)
				ret := expr.BeanExpr(field.Tag.Get("inject"), this.ExprMap)
				if ret != nil && !ret.IsEmpty() {
					retValue := ret[0]
					if retValue != nil {
						v.Field(i).Set(reflect.ValueOf(retValue)) //给结构体成员设置从ioc容器中取出的值
						this.Apply(retValue)                      //递归调用，因为可能会出现嵌套依赖的情况
					}
				}
			} else { //单例模式(传统模式)
				if get_v := this.Get(field.Type); get_v != nil { //从beanMapper中取出对应的值
					v.Field(i).Set(reflect.ValueOf(get_v)) //给结构体成员设置从ioc容器中取出的值
					this.Apply(get_v)                      //递归调用，因为可能会出现嵌套依赖的情况
				}
			}
		}
	}
}

// 向ioc容器中添加依赖
func (this *BeanFactoryImpl) Config(cfgs ...interface{}) {
	for _, cfg := range cfgs {
		t := reflect.TypeOf(cfg)
		if t.Kind() != reflect.Ptr {
			panic("required ptr object") //必须是指针对象
		}
		if t.Elem().Kind() != reflect.Struct {
			continue
		}

		this.Set(cfg)                       //把config本身也加入bean
		this.ExprMap[t.Elem().Name()] = cfg //自动构建 ExprMap
		this.Apply(cfg)                     //处理依赖注入 (new)

		v := reflect.ValueOf(cfg)
		//遍历所有的类方法
		for i := 0; i < t.NumMethod(); i++ {
			method := v.Method(i)
			callRet := method.Call(nil) //执行类方法

			if callRet != nil && len(callRet) == 1 {
				this.Set(callRet[0].Interface()) //将类方法的执行结果塞入到ioc容器中
			}
		}
	}
}
