package main

import (
	"log"
	"math"
)

type FuncDelta func(x, y float64) float64 // 声明了一个函数类型

func main() {

	log.SetFlags(log.Lshortfile)

	val := 3.0

	// 牛顿法求解 x 平方根
	log.Printf("%v的2次方根 is %v \n", val, Sqrt(val, NewtonDelta2))
	log.Printf("golang 标准库平方根结果: %v \n ", math.Sqrt(val))

	// 牛顿法求解 x 3次方根
	log.Printf("%v的3次方根 is %v \n", val, Sqrt(val, NewtonDelta3))

}

// 牛顿法求解 x 方根
func Sqrt(val float64, delta FuncDelta) float64 {

	const step int = 5

	x := val / 2 // 初值 x0

	for i := 0; i < step; i++ {

		log.Printf("step%v:%v \n", i, x)

		delta := delta(x, val)
		x = x - delta
	}

	return x
}

// 牛顿法2次方根算子
func NewtonDelta2(x, y float64) float64 {
	return (x*x - y) / (2 * x)
}

// 牛顿法3次方根算子
func NewtonDelta3(x, y float64) float64 {
	return (x*x*x - y) / (3 * x * x)
}
