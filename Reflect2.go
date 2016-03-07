/*
反射reflection
1.	反射可大大提高程序的灵活性，使得interface{}有更大的发挥余地
2.	反射使用TypeOf和ValueOf函数从接口中获取目标对象信息
3.	反射会将匿名字段作为独立字段(匿名字段本质)
4.	想要利用反射修改对象状态，前提是interface.data是setable，即pointer-interface
5.	通过反射可以“动态”调用方法
*/

package main

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

func main() {
	/*
		m := Manager{User: User{1, "OK", 25}, title: "123"}
		t := reflect.TypeOf(m)

		//fmt.Printf("%#v\n", t.Field(0))
		//Output:reflect.StructField{Name:"User", PkgPath:"", Type:(*reflect.rtype)(0x80c9b00), Tag:"", Offset:0x0, Index:[]int{0}, Anonymous:true}

		//fmt.Printf("%#v\n", t.FieldByIndex([]int{0, 0}))
		//Output:reflect.StructField{Name:"Id", PkgPath:"",	Type:(*reflect.rtype)(0x80b7ce0), Tag:"", Offset:0x0, Index:[]int{0}, Anonymous:false}
	*/

	/*
		x := 123
		v := reflect.ValueOf(&x)
		v.Elem().SetInt(999)
		fmt.Println(x)

		u := User{1, "OK", 22}
		Set(&u)
		fmt.Println(u)
	*/

	//利用反射进行动态调用
	u := User{1, "OK", 22}
	v := reflect.ValueOf(u)
	mv := v.MethodByName("Hello")
	args := []reflect.Value{reflect.ValueOf("Joe")}
	mv.Call(args)
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
