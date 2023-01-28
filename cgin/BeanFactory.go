package cgin

import "reflect"

type BeanFactory struct {
	beans []interface{}
}

func NewBeanFactory() *BeanFactory {
	bf := &BeanFactory{beans: make([]interface{}, 0)}
	bf.beans = append(bf.beans, bf)
	return bf
}

// 往内存中塞入bean
func (this *BeanFactory) setBean(beans ...interface{}) {
	this.beans = append(this.beans, beans...)
}

// 外部使用
func (this *BeanFactory) GetBean(bean interface{}) interface{} {
	return this.getBean(reflect.TypeOf(bean))
}

// 得到 内存中预先设置好的bean对象
func (this *BeanFactory) getBean(t reflect.Type) interface{} {
	for _, p := range this.beans {
		if t == reflect.TypeOf(p) {
			return p
		}
	}
	return nil
}

// 给外部用的 （后面还要改,这个方法不处理注解)
func (this *BeanFactory) Inject(object interface{}) {
	vObject := reflect.ValueOf(object)
	if vObject.Kind() == reflect.Ptr { //由于不是控制器 ，所以传过来的值 不一定是指针。因此要做判断
		vObject = vObject.Elem()
	}
	for i := 0; i < vObject.NumField(); i++ {
		f := vObject.Field(i)
		if f.Kind() != reflect.Ptr || !f.IsNil() {
			continue
		}
		if p := this.getBean(f.Type()); p != nil && f.CanInterface() {
			f.Set(reflect.New(f.Type().Elem()))
			f.Elem().Set(reflect.ValueOf(p).Elem())
		}
	}
}

// 把bean注入到控制器中 (内部方法,用户控制器注入。并同时处理注解)
func (this *BeanFactory) inject(class IClass) {
	vClass := reflect.ValueOf(class).Elem() //接口是一个指针，所以要用Elem()获取指针指向的值，比如UserClass{}
	vClassT := reflect.TypeOf(class).Elem()
	for i := 0; i < vClass.NumField(); i++ { //遍历结构体的成员
		f := vClass.Field(i)                       //得到结构体成员
		if f.Kind() != reflect.Ptr || !f.IsNil() { //如果成员不为空，且不是指针类型
			continue //就继续取下一个成员
		}
		if IsAnnotation(f.Type()) { //代表是注解 ,则单独处理
			f.Set(reflect.New(f.Type().Elem()))
			f.Interface().(Annotation).SetTag(vClassT.Field(i).Tag)
			this.Inject(f.Interface())
			continue
		}
		if p := this.getBean(f.Type()); p != nil { //否则就判断成员的类型是否是在数据库连接对象中的一个，是则返回数据库连接对象
			f.Set(reflect.New(f.Type().Elem()))     //reflect.New(f.Type().Elem())是创建一个f.Type().Elem()类型的结构体对象，之后使用set给f赋值
			f.Elem().Set(reflect.ValueOf(p).Elem()) //给f设置上数据库连接对象的value
		}
	}
}
