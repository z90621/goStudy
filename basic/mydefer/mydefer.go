package main

import (
	"fmt"
)

func main() {
	fmt.Println("=========================")
	fmt.Println("return:", fun1())

	fmt.Println("=========================")
	fmt.Println("return:", fun2())
	fmt.Println("=========================")

	fmt.Println("return:", fun3())
	fmt.Println("=========================")

	fmt.Println("return:", fun4())
}

// 有名返回值是如何返回的呢？return并不是原子性的
// 1.return如果有计算进行计算
// 2.把值赋值给有名返回变量
// 3.计算defer
// 4.返回
func fun1() (i int) {
	println("当前i为：", i)
	defer func() {
		i++
		fmt.Println("defer2:", i) // 打印结果为 defer2: 2
	}()

	// 规则二 defer执行顺序为先进后出

	defer func() {
		i++
		fmt.Println("defer1:", i) // 打印结果为 defer1: 1
	}()

	// 规则三 defer可以读取有名返回值（函数指定了返回参数名）
	println("fun1准备return")
	return i + 2 //这里实际结果为2。如果是return 100呢
}

// 匿名返回值
func fun2() int {
	var i int = 10
	defer func() {
		i++
		fmt.Println("defer2:", i) // 打印结果为 defer2: 2
	}()

	defer func() {
		i++
		fmt.Println("defer1:", i) // 打印结果为 defer1: 1
	}()
	return i //10
}

func fun3() (r int) {
	t := 5
	defer func() {
		t = t + 5
		fmt.Println(t)
	}()
	return t
}

func fun4() int {
	i := 8
	// 规则一 当defer被声明时，其参数就会被实时解析
	defer func(i int) {
		i = 99
		fmt.Println(i)
	}(i)
	i = 19
	return i
}
