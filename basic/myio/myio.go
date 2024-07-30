package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

/**
  os os.file实现io reader writer接口
  io
  io/ioutil 内部调用os
  bufio 组合io，使用buf
  file.Read > ioutil >bufio
**/

// O_RDONLY int = syscall.O_RDONLY // 只读打开文件和os.Open()同义
// O_WRONLY int = syscall.O_WRONLY // 只写打开文件
// O_RDWR   int = syscall.O_RDWR   // 读写方式打开文件
// O_APPEND int = syscall.O_APPEND // 当写的时候使用追加模式到文件末尾
// O_CREATE int = syscall.O_CREAT  // 如果文件不存在，此案创建
// O_EXCL   int = syscall.O_EXCL   // 和O_CREATE一起使用，只有当文件不存在时才创建
// O_SYNC   int = syscall.O_SYNC   // 以同步I/O方式打开文件，直接写入硬盘
// O_TRUNC  int = syscall.O_TRUNC  // 如果可以的话，当打开文件时先清空文件

func _bufio() {
	f, _ := os.OpenFile("tmp.txt", os.O_CREATE|os.O_RDWR, 0666)
	defer f.Close()
	rd := bufio.NewReader(f)
	for {
		bs, e := rd.ReadSlice('-')
		if e == io.EOF {
			break
		}
		fmt.Println(string(bs))
	}
	// f.Seek(0, 0)
	wr := bufio.NewWriter(f)
	wr.WriteString("123-321-你好-ad")
	wr.Flush()
}

func _ioutil() {
	//ioutil一大堆方法在1.6后废弃，推荐使用os内同名方法
}

func main() {
	// os.File
	// f, _ := os.OpenFile("xxxx", os.O_RDWR, 0666)
	// f.Read()
	// io.ByteReader
	_bufio()
}
