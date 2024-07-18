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
