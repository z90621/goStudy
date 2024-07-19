package myfunction

import (
	"fmt"
)

var G int = 0

func F1() {
	y := func() int {
		fmt.Printf("G: %d, G的地址:%p\n", G, &G)
		G += 1
		return G
	}
	fmt.Println(y(), y)
	fmt.Println(y(), y)
	fmt.Println(y(), y) //y的地址

	// 影响全局变量G，注意z的匿名函数是直接执行，所以结果不变
	//z存储了匿名函数的结果
	z := func() int {
		G += 1
		println("----G=", G)
		return G
	}()
	fmt.Println(z, &z)
	fmt.Println(z, &z)
	fmt.Println(z, &z)
	fmt.Printf("z type is %T", z) //is int

}

func Bibao2() {
	vf := f()
	fmt.Println(vf(1), &vf)
	fmt.Println(vf(1), &vf)
	fmt.Println(vf(1), &vf)
	fmt.Println("-----------------")
	vf2 := f() //新环境
	fmt.Println(vf2(1), &vf2)
}

func f() func(int) int {
	var i int //地址不会变化，环境+函数 = 闭包
	return func(d int) int {
		fmt.Printf("i: %d, i的地址:%p\n", i, &i)
		i += d
		return i
	}
}
