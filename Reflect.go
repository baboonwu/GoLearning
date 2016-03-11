/*
反射reflection
1.	反射可大大提高程序的灵活性，使得interface{}有更大的发挥余地
2.	反射使用TypeOf和ValueOf函数从接口中获取目标对象信息
3.	反射会将匿名字段作为独立字段(匿名字段本质)
4.	想要利用反射修改对象状态，前提是interface.data是setable，即pointer-interface
5.	通过反射可以“动态”调用方法
*/

package GoLearning

import (
	"fmt"
	"reflect"
)

type User struct {
	Id   int
	Name string
	Age  int
}

type Manager struct {
	User
	title string
}

func (u User) Hello(name string) {
	fmt.Println("Hello", name, ",My name is", u.Name)
}

func Info(o interface{}) {
	t := reflect.TypeOf(o)
	fmt.Println("Type:", t.Name())

	if k := t.Kind(); k != reflect.Struct { //若为non-struct类型，则返回
		fmt.Println("XXX")
		return
	}

	v := reflect.ValueOf(o)
	fmt.Println("Fields:")

	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		val := v.Field(i).Interface()
		fmt.Printf("%6s: %v = %v\n", f.Name, f.Type, val)
	}

	for i := 0; i < t.NumMethod(); i++ {
		m := t.Method(i)
		fmt.Printf("%6s: %v\n", m.Name, m.Type)
	}
}

func Set(o interface{}) {
	v := reflect.ValueOf(o)

	if v.Kind() == reflect.Ptr && !v.Elem().CanSet() {
		fmt.Println("XXX")
		return
	} else {
		v = v.Elem()
	}

	f := v.FieldByName("Name")
	if !f.IsValid() {
		fmt.Println("BAD")
		return
	}

	if f.Kind() == reflect.String {
		f.SetString("ByeBye")
	}
}
