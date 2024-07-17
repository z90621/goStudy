package myarray

import (
	"fmt"
	"reflect"
)

//数组长度也是数组类型的一部分，所以[5]int和[10]int是属于不同类型的

func Declar() []int {
	var arrAge = [5]int{18, 20, 15, 22, 16}
	fmt.Printf("type is %T\n", arrAge) //长度也是类型要素之一
	var arrAge2 = new([5]int)
	fmt.Printf("arrAge2 type is %T\n", arrAge2)
	var arrName = [5]string{3: "Chris", 4: "Ron"} //指定索引位置初始化
	println(arrName[3])
	// {"","","","Chris","Ron"}
	var arrCount = [4]int{500, 2: 100} //指定索引位置初始化 {500,0,100,0}
	println(arrCount[0])
	var arrLazy = [...]int{5, 6, 7, 8, 22} //数组长度初始化时根据元素多少确定
	println(len(arrLazy))
	var arrPack = [...]int{10, 5: 100} //指定索引位置初始化，数组长度与此有关 {10,0,0,0,0,100}
	println(arrPack[5])
	// var arrRoom [20]int
	// var arrBed = new([20]int)
	return nil
}

// slice
func Slice_() {
	//切割数组生成slice
	var arrAge = [5]int{18, 20, 15, 22, 16}
	var s1 []int = arrAge[0:2:5] //2-0长度，5-0容量
	s1[0] = 1
	println(arrAge[0]) // == 1

	arrName := []string{"Tom"}      //类似数组但不指定长度 slice
	var arrName2 = [1]string{"Tom"} //array
	println(arrName[0])
	fmt.Printf("arrName len is %d,cap is %d,type is %T\n", len(arrName), cap(arrName), arrName)
	fmt.Printf("arrName2 len is %d,cap is %d,type is %T\n", len(arrName2), cap(arrName2), arrName2)
	println(reflect.TypeOf(arrName).Kind() == reflect.Slice)
	println(reflect.TypeOf(arrName2).Kind())
	//如果未定义数组，var slice1 []type = make([]type, len,cap)
	noArray := make([]string, 10, 20)
	noArray[0] = "one"
	//索引不能超过长度，如果想使用剩余容量，可以对这个切片进行切片重组，切大点
	//重组后会继续引用之前的数组，如果是从一大块切成一小块使用会导致空间浪费
	// noArray[20] = "ww"
	sli := make([]int, 5, 10)
	fmt.Printf("切片sli长度和容量：%d, %d\n", len(sli), cap(sli))
	sli[0] = 1
	fmt.Println(sli)
	newsli := sli[:cap(sli)]
	newsli[9] = 9
	fmt.Println(newsli)
	//切片重组，重组小切片推荐使用copy()
	bigS := make([]int, 10, 1000)
	bigS[0] = 1
	smallS := make([]int, 3)
	copy(smallS, bigS[0:3]) // bigS[0:3]会被内存释放
	fmt.Println(smallS)

	//append append()函数将 0 个或多 z个具有相同类型S的元素追加到切片s后面并且返回新的切片
	//append()函数操作后，有没有生成新的切片需要看原有切片的容量是否足够。
	ap1 := make([]int, 2, 3)
	ap1[0] = 1
	ap1[1] = 2
	fmt.Println(ap1)
	ap2 := ap1[1:]
	ap2[0] = 4
	fmt.Println(ap1)
	fmt.Println(ap2)
	ap2 = append(ap2, 23)
	ap2[0] = 6
	fmt.Println(ap1)
	fmt.Println(ap2)
	ap2 = append(ap2, 24)
	//从这开始超出原本数据长度，扩容到新数组返回，修改不再影响原来切片
	ap2[0] = 8
	fmt.Println(ap1)
	fmt.Println(ap2)

}
