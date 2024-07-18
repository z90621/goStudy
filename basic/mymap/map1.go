package mymap

import (
	"fmt"
)

func MyMap() {

	// var m1 map[string]float32 -->未初始化添加会panic，slice不会
	m1 := make(map[string]int, 1)
	m1["one"] = 1
	fmt.Println(m1)
	//判断值存在否
	if _, ok := m1["two"]; !ok {
		println("m2不存在key[two]")
	}
	delete(m1, "three") //无值也不会报错

	//range值为拷贝，如果要修改原集合数据，可以使用下标更新
	for k, v := range m1 {
		println("key is", k, "val is", v)
		println("准备直接修改val")
		v = 2
		println(m1["one"])
		println("准备使用下标修改val")
		m1[k] = 3
		println(m1["one"])
	}

}
