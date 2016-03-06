/*
接口Interface
1.	接口是一个或多个方法签名的集合
2.	只要某个类型拥有该接口的所有方法签名，即算实现该接口，
	无需显示声明实现了哪个接口，这称为Structural Typing
3.	接口只有方法声明，没有实现，没有数据字段
4.	接口可以匿名嵌入其他接口，或嵌入到结构中
5.	将对象赋值给接口时，会发生拷贝，而接口内部存储的是指向这个
	复制品的指针，既无法修改复制品的状态，也无法获取指针
6.	只有当接口存储的类型和对象都为nil时，接口才等于nil
7.	接口调用不会做receive的自动转换
8.	接口同样支持匿名字段方法
9.	接口也同样可实现类似OOP中的多态
10.	空接口可以作为任何类型数据的容器
*/
/*interface是一组method的组合，我们通过interface来定义对象的一组行为。
interface类型定义了一组方法，如果某个对象实现了某个接口的所有方法，
则此对象就实现了此接口，不需要显示声明。*/

package main

import (
	"fmt"
)

//10.空接口可以作为任何类型数据的容器
type empty interface {
}

//4.接口可以匿名嵌入其他接口，或嵌入到结构中
type USB interface { //USB可以转换为Connector:因为USB嵌入了Connector
	Name() string
	Connector
}

type Connector interface { //Connector不能转换为USB
	Connect()
}

type PhoneConnector struct {
	name string
}

func (pc PhoneConnector) Name() string {
	return pc.name
}

func (pc PhoneConnector) Connect() {
	fmt.Println("Connected:", pc.name)
}

func Disconnect(usb interface{}) {
	/*if pc, ok := usb.(PhoneConnector); ok {//类型判断
		fmt.Println("Disconnected:", pc.name)
		return
	}
	fmt.Println("Unknown device.")*/

	switch v := usb.(type) { //类型判断：系统自动判断
	case PhoneConnector:
		fmt.Println("Disconnected:", v.name)
	default:
		fmt.Println("Unknown device.")
	}
}

type TVConnector struct {
	name string
}

func (tv TVConnector) Connect() {
	fmt.Println("Connect:", tv.name)
}

func main() {
	/*a := PhoneConnector{"PhoneConnector"}
	a.Connect()
	Disconnect(a)*/

	/*pc := PhoneConnector{"PhoneConnector"}
	var a Connector
	a = Connector(pc)
	a.Connect()
	Disconnect(a)*/

	/*tv := TVConnector("TVConnector")
	var a USB
	a = USB(tv) //cannot convert tv (type TVConnector) to type USB:
	//TVConnector does not implement USB (missing Name method)
	a.Connect()*/

	/*5.将对象赋值给接口时，会发生拷贝，而接口内部存储的是指向这个
	复制品的指针，既无法修改复制品的状态，也无法获取指针*/
	/*pc := PhoneConnector{"PhoneConnector"}
	var a Connector
	a = Connector(pc)
	a.Connect()
	pc.name = "pc"
	a.Connect()*/

	//6.只有当接口存储的类型和对象都为nil时，接口才等于nil
	var e interface{}
	fmt.Println(e == nil) //true
	var p *int = nil
	e = p
	fmt.Println(e == nil) //false
}
