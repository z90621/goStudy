package mypointer

import (
	"fmt"
	"log"
	"runtime"
	"time"
)

func Point() {
	var a int = 1
	// var pt1 *int
	var pt2 *int = &a

	fmt.Println(pt2)
	fmt.Println(&pt2)
	fmt.Println(*pt2)
}

type Person struct {
	Name string
	Age  int
}

func (p *Person) Close() {
	p.Name = "NewName"
	log.Println(p)
	log.Println("Close")
}

func (p *Person) NewOpen() {
	log.Println("Init")
	//SetFinalizer 函数是 Go 语言的运行时包 runtime 提供的一个功能，
	//它允许你设置一个函数，这个函数会在垃圾收集器（GC）确定对象 obj 不再可达时被调用，
	//以执行清理工作。
	runtime.SetFinalizer(p, (*Person).Close)
}

func Tt(p *Person) {
	p.Name = "NewName"
	log.Println(p)
	log.Println("Tt")
}

// 查看内存情况
func Mem(m *runtime.MemStats) {
	runtime.ReadMemStats(m)
	log.Printf("%d Kb\n", m.Alloc/1024)
}

func Main() {
	var m runtime.MemStats
	Mem(&m)

	var p *Person = &Person{Name: "lee", Age: 4}
	p.NewOpen()
	log.Println("Gc完成第一次")
	log.Println("p:", p)
	runtime.GC()
	time.Sleep(time.Second * 5)
	Mem(&m)

	var p1 *Person = &Person{Name: "Goo", Age: 9}
	runtime.SetFinalizer(p1, Tt)
	log.Println("Gc完成第二次")
	time.Sleep(time.Second * 2)
	runtime.GC()
	time.Sleep(time.Second * 2)
	Mem(&m)

}
