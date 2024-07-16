package mystring

import (
	"bytes"
	"fmt"
	"strings"

	// "strconv"
	"unicode/utf8"
	// "unicode"
)

func MyString() {
	/**
		strconv包提供了布尔型、整型数、浮点数和对应字符串的相互转换，还提供了双引号转义相关的转换。

		unicode包提供了IsDigit、IsLetter、IsUpper和IsLower等类似功能，它们用于给字符分类。

		strings 包提供了很多操作字符串的简单函数，通常一般的字符串操作需求都可以在这个包中找到
	**/
	str := "hi"
	var buffer bytes.Buffer
	var builder strings.Builder
	buffer.WriteString(str)
	buffer.WriteString(" second write")
	builder.WriteString(buffer.String())
	builder.WriteString("汉字写入")
	str2 := builder.String()
	for i := 0; i < len(str2); i++ {
		fmt.Println(str2[i])
	}
	fmt.Println()
	//汉字的遍历
	for k, v := range str2 {
		fmt.Printf("k is %d,v is %c\n", k, v)
	}
	fmt.Printf("str2 len is %d\n", len(str2))
	fmt.Printf("str2 true len is %d\n", utf8.RuneCountInString(str2))
}
