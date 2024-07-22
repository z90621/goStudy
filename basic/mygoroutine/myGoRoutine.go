package mygoroutine

//Go 语言在语言层面上支持了并发，goroutine是Go语言提供的一种用户态线程，有时我们也称之为协程
// 调度器的主要有4个重要部分，分别是M、G、P、Sched。
// M (work thread) 代表了系统线程内核线程，由操作系统管理。
// P (processor) 衔接M和G的调度上下文，它负责将等待执行的G与M对接。P的数量可以通过GOMAXPROCS()来设置，它其实也就代表了真正的并发度，即有多少个goroutine可以同时运行。
// G (goroutine) 协程的实体，包括了调用栈，重要的调度信息，例如channel等。

/**runtime.NumCPU()        // 返回当前CPU内核数
runtime.GOMAXPROCS(2)  // 设置运行时最大可执行CPU数
runtime.NumGoroutine() // 当前正在运行的协程 数
**/

//golang调度器设计策略
/**
复用线程:work stealing 空闲线程本地队列为空，全局也为空，那么窃取其他线程的任务队列任务执行，
						双端队列，窃取时从尾部窃取
		:hand off 如果当前G获取资源发生系统调用，当前M0会将原本P队列移交给其他M1执行，
		当调用完成，M0会看有没有空闲P，有的话拿到P并进行G处理，如果没有，将G丢到全局队列
		等待其他P来拿
利用并行
抢占
全局G队列
**/

import (
	"fmt"
	"sync"
)

func SomeApi() {
	//等待组
	var wg sync.WaitGroup
	//only run once
	var once sync.Once
	wg.Add(2)

	go func() {
		fmt.Println("hi")
		wg.Done()
	}()
	go func() {
		for i := 1; i < 10; i++ {
			once.Do(func() { fmt.Println("hello") })
		}
		wg.Done()
	}()
	wg.Wait()

}

func _syncMap() {
	// var m sync.Map
	/**
	store
	loadOrStore key 存在 ->返回值，不存在，存入
	load
	delete

	**/
}
