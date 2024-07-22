package mystruct

import (
	"fmt"
	"reflect"
)

type animal struct {
	name   string "名字"
	danger int    `json:"危险程度"` //结构体标签，可以通过reflect获取
}

func (a *animal) ToString() {
	println("animal named", a.name, "and danger is", a.danger)
}

func (a *animal) SetName(newName string) {
	a.name = newName
	println(a.name)
}

func (a animal) SetName2(newName string) {
	fmt.Printf("a address is %p\n", &a)
	fmt.Printf("setName2 Object type is %T\n", a)
	a.name = newName
}

type myAnimal animal

type mmc struct {
	animal
}

type action interface {
	eat()
	drink()
}

func (a *animal) eat() {
	println("eat")
}
func (a animal) drink() {
	println("drink")
}

func TAnimal() {
	my := myAnimal{"12", 1}
	// my.setName2("1")
	// an := new(animal)
	// an.setName2("!")
	// mc := mmc{animal{"l;", 1}}
	// mc.setName2("1")
	// mc.toString()

	an2 := mmc{animal{"asas", 2}}
	//下面这个声明编译无法通过，如果一个类型实现一个接口以指针方法来实现，那么只有指针类型能实现这个接口
	// var an3 action = an2
	//下面这个声明正确，如果一个类型实现接口以值方法实现，不管你是值还是指针变量都能实现这个接口
	// var an4 action = &an2
	//组合继承与实现接口异曲同工，也是你指针类型变量才能涵盖所有方法实现
	//从这里我们可以很清晰的看成，当以组合的方式实现继承时，值类型只能调用嵌入对象的值类型方法
	//指针类型则可以调用全部，之所以值类型变量依然可以进行调用指针类型嵌入方法，这里是go做了语法糖
	methodSet(an2)
	methodSet(&an2)
	//自定义新结构体
	methodSet(my)

}

func methodSet(a interface{}) {
	t := reflect.TypeOf(a)
	fmt.Printf("%T\n", a)
	for i, n := 0, t.NumMethod(); i < n; i++ {
		m := t.Method(i)
		fmt.Println(i, ":", m.Name, m.Type)
	}
}
