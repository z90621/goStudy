package myerror

import (
	"fmt"
)

func MyError() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("捕捉到异常%s\n", r)
		}
	}()

	if x := -1; x < 0 {
		panic("x 小于0")
	}

}

func MyDefer() {
	i := 1
	//1.defer参数会在开始便初始化，而不会因为最后执行改变参数
	// 	参数评估时机：
	// defer语句中的参数（如果有的话）在defer语句执行时进行评估，
	//而不是在defer调用时。这意味着参数的值在defer语句声明时就确定了
	defer func() int { //无参数 打印2
		println("2->", i)
		return i
	}()
	//以下参数维持初始化状态也就是数值1进行操作
	defer fmt.Println("result =>", func() int { println("in result->", i); return i * 2 }())
	defer fmt.Println("3->", i*2)
	defer fmt.Println("4->", i)
	defer println("5->", i)

	i++
	//以上执行顺序结果为
	// in result-> 1 初始化参数执行了内部打印
	// 5-> 1
	// 4-> 1
	// 3-> 2
	// result
	// 2-> 2 剩下的都遵守入出栈执行顺序，
	//区别在于除最后一个，defer打印的函数带有参数在正常执行便固定为1
}
