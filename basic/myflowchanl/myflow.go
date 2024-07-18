package myflow

import (
	"fmt"
	"time"
)

func MyFlow() {
	animal := "pig"
	switch animal {
	case "duck":
		println("鸭子")
	case "pig", "wide pig": //多个条件
		println("^(*￣(oo)￣)^")
	default:
		println("default")

	}

	switch { //无变量自动执行分支为true的
	case "1" == "1":
		println("yes,obs 1=1")
		fallthrough //强行执行分支代码，不用判断.穿透一层
	case "1" == "2":
		println("no")
		fallthrough
	case "2" == "2":
		println("yes is 2")
	default:
		println("default")
	}

}

func MySelect() {
	//channel 也需要初始化才能添加值
	var ch1 chan int = make(chan int)
	//单个协程内写以下代码形成死锁，无缓冲chan的读写是同时发生的，单个无法完成，同步过程
	// ch1 <- 1
	// <-ch1
	go func(ch chan int) {
		ch <- 1
	}(ch1)
	x := <-ch1
	println("读取数据：", x)
	//关闭通道
	close(ch1)
	//关闭后再往其发送数据将会panic,可以通过拿数据判断
	if _, ok := <-ch1; !ok {
		println("ch is close")
	}

	//带缓冲区的channel不会deadlock
	ch2 := make(chan int, 2)
	ch2 <- 100
	println(<-ch2)
}

func MyChannel() {
	//使用空结构体通信，不占内存
	send := make(chan struct{})
	recive := make(chan struct{})

	//<-限定chan为输入通道还是输出通道
	go func(s chan<- struct{}) {
		s <- struct{}{}
		recive <- struct{}{}
	}(send)

	go func(r <-chan struct{}) {
		x := <-r
		fmt.Println(x)
	}(recive)

}

type filed struct {
	name string
}

func (e *filed) print() {
	println(e.name)
	e.name = "123"
	println(e.name)
}

func MyForRange() {
	// s := make([]any, 10)

	// arr := []string{"123", "12", "1"}

	// for i, v := range arr {
	// 	println("i=", i, "v=", v)
	// }
	fileds := []filed{{name: "zs"}, {"ls"}, {"ww"}}
	for _, v := range fileds {
		v.print()
	}
	fmt.Println(fileds)

	time.Sleep(3 * time.Second)

}

// range获取的临时变量地址不相同了，
// 虽然还是数值值的拷贝，修改不会影响原来，但貌似不会出现重复利用地址导致最后值全相等的情况
func MyTest() {
	arr1 := []int{1, 2, 3}

	arr2 := make([]*int, 0)

	for _, v := range arr1 {

		arr2 = append(arr2, &v)
	}

	for _, v := range arr2 {
		if v == nil {
			continue
		}
		fmt.Println(v)
		fmt.Println(*v)
	}
}
