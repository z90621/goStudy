package mystruct

import (
	"fmt"
)

type animal struct {
	name   string "名字"
	danger int    `json:"危险程度"` //结构体标签，可以通过reflect获取
}

func (a *animal) toString() {
	println("animal named", a.name, "and danger is", a.danger)
}

func (a *animal) setName(newName string) {
	a.name = newName
	println(a.name)
}

func (a animal) setName2(newName string) {
	fmt.Printf("a address is %p\n", &a)
	fmt.Printf("setName2 Object type is %T\n", a)
	a.name = newName
}

func TAnimal() {
	an := new(animal)
	an.name = "pig"
	an.danger = 1
	println(&an)
	fmt.Printf("an type is %T\n", an)
	an.setName2("zz")
	an.toString()
}
