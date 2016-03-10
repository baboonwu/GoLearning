package GoLearning

import (
	"fmt"
	"log"
	"reflect"
	"testing"
)

func init() {

	log.SetFlags(log.Lshortfile)
}

func Test_Interface1(t *testing.T) {

	a := PhoneConnector{"PhoneConnector"}
	a.Connect()
	Disconnect(a)

	/*log.Println(result, errResult, e)

	if e != nil {
		t.Error(e)
	}*/
}

func Test_Interface2(t *testing.T) {

	pc := PhoneConnector{"PhoneConnector"}
	var a Connector
	a = Connector(pc)
	a.Connect()
	Disconnect(a)
}

func Test_Interface3(t *testing.T) {

	/*5.将对象赋值给接口时，会发生拷贝，而接口内部存储的是指向这个
	复制品的指针，既无法修改复制品的状态，也无法获取指针*/
	pc := PhoneConnector{"PhoneConnector"}
	var a Connector
	a = Connector(pc)
	a.Connect()
	pc.name = "pc"
	a.Connect()
}

func Test_Interface4(t *testing.T) {

	//6.只有当接口存储的类型和对象都为nil时，接口才等于nil
	var e interface{}
	fmt.Println(e == nil) //true
	var p *int = nil
	e = p
	fmt.Println(e == nil) //false
}

func Test_Reflect1(t *testing.T) {

	u := User{1, "Joe", 25}
	Info(&u)
}

func Test_Reflect2(t *testing.T) {

	m := Manager{User: User{1, "OK", 25}, title: "123"}
	tp := reflect.TypeOf(m)

	fmt.Printf("%#v\n", tp.Field(0))
	//Output:reflect.StructField{Name:"User", PkgPath:"", Type:(*reflect.rtype)(0x80c9b00), Tag:"", Offset:0x0, Index:[]int{0}, Anonymous:true}

	fmt.Printf("%#v\n", tp.FieldByIndex([]int{0, 0}))
	//Output:reflect.StructField{Name:"Id", PkgPath:"",	Type:(*reflect.rtype)(0x80b7ce0), Tag:"", Offset:0x0, Index:[]int{0}, Anonymous:false}

}

func Test_Reflect3(t *testing.T) {

	x := 123
	v := reflect.ValueOf(&x)
	v.Elem().SetInt(999)
	fmt.Println(x)

	u := User{1, "OK", 22}
	Set(&u)
	fmt.Println(u)
}

func Test_Reflect4(t *testing.T) {

	//利用反射进行动态调用
	u := User{1, "OK", 22}
	v := reflect.ValueOf(u)
	mv := v.MethodByName("Hello")
	args := []reflect.Value{reflect.ValueOf("Joe")}
	mv.Call(args)
}
