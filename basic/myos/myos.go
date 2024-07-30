package main

import (
	"fmt"
	"os"
	"os/signal"
	"os/user"
	"syscall"
)

/**
os
 --exce
 --signal
 --user

**/

// 注意操作系统的差别
func _user() {
	// 获取当前用户信息
	currentUser, err := user.Current()
	if err != nil {
		fmt.Printf("Error getting current user: %v\n", err)
		return
	}

	// 打印当前用户信息
	fmt.Printf("Username: %s\n", currentUser.Username)
	fmt.Printf("User ID: %d\n", currentUser.Uid)
	fmt.Printf("Group ID: %d\n", currentUser.Gid)
	fmt.Printf("Home directory: %s\n", currentUser.HomeDir)
	fmt.Printf("Name: %s\n", currentUser.Name)
	fmt.Println(currentUser.GroupIds())
}

/*
*
信号处理通常用于优雅地关闭程序或执行清理任务。
在处理信号时，应该小心操作，因为信号处理函数可能会在任何时候中断程序执行。
*
*/
func _signal() {
	ch := make(chan os.Signal, 10)
	signal.Notify(ch, os.Interrupt, os.Kill)
	sig := <-ch
	fmt.Printf("sig: %v\n", sig)
}

// 启动外部系统命令和二进制可执行文件
// 第一个参数是要运行的进程，
// 第二个参数用来传递选项或参数，
// 第三个参数是含有系统环境基本信息的结构体。
func startProcess() {
	// 设置程序路径和参数
	prog := `D:\Notepad++\notepad++.exe`
	args := []string{"notepad.exe"}

	// 设置进程属性
	env := os.Environ()  // 获取当前环境变量
	dir, _ := os.Getwd() // 获取当前工作目录
	attr := &os.ProcAttr{
		Env: env,
		Dir: dir,
		Files: []*os.File{
			os.Stdin,
			os.Stdout,
			os.Stderr,
		},
		Sys: &syscall.SysProcAttr{},
	}

	// 启动进程
	process, err := os.StartProcess(prog, args, attr)
	if err != nil {
		panic(err)
	}

	// 等待进程结束
	process.Wait()
	println("notpad exit")
}

func doc() {
	tmp := "tmp.txt"
	f, _ := os.Create(tmp)
	defer f.Close()
	// Stat返回描述文件f的FileInfo类型值
	fmt.Println(f.Stat())
	f.WriteString("fk")
	//f.Sync() //在缓冲区挂起的数据进行稳定的存储-写入磁盘
	f.Seek(0, 0)
	b := make([]byte, 64)
	var i int
	for {
		i, _ = f.Read(b)
		if i == 0 {
			break
		}

	}
	str := string(b)
	fmt.Printf("string(b): %v\n", str)

}

func main() {
	// fmt.Println(os.Environ())
	// fmt.Println(os.Args)
	// startProcess()
	// _signal()
	_user()
}
