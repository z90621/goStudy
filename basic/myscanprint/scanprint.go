package main

//在fmt包，有关格式化输入输出的方法就两大类：Scan 和 Print ，分别在scan.go 和 print.go 文件中。
import "fmt"

/**
Println、Fprintln、Sprintln  输出内容时会加上换行符；
Print、Fprint、Sprint        输出内容时不加上换行符；
Printf、Fprintf、Sprintf     按照指定格式化文本输出内容。

Print、Printf、Println      输出内容到标准输出os.Stdout；
Fprint、Fprintf、Fprintln   输出内容到指定的io.Writer；
Sprint、Sprintf、Sprintln   输出内容到字符串。

Scanln、Fscanln、Sscanln    读取到换行时停止，并要求一次提供一行所有条目；
Scan、Fscan、Sscan          读取内容时不关注换行；
Scanf、Fscanf、Sscanf       根据格式化文本读取。

Scan、Scanf、Scanln     从标准输入os.Stdin读取文本；
Fscan、Fscanf、Fscanln  从指定的io.Reader接口读取文本；
Sscan、Sscanf、Sscanln  从一个参数字符串读取文本。
**/

func _sprintf() {
	print(fmt.Sprintf("你好，你在%d", 1))
}

func _scan() {
	var name string
	var age int
	fmt.Scan(&name)
	fmt.Scan(&age)
	fmt.Printf("name is %v,age is %d", name, age)
}

func _sscan() {
	// 定义一个字符串，其中包含格式化的输入数据
	input := "Hello, World! 123"

	// 使用 fmt.Sscanf 函数解析输入
	var (
		greeting string
		number   int
	)

	// 格式化输入的字符串，其中 %s 用于字符串，%d 用于整数
	format := "Hello, %s! %d"

	// 使用 fmt.Sscanf 解析输入
	fmt.Sscanf(input, format, &greeting, &number)

	// 打印解析后的数据
	fmt.Printf("Greeting: %s\n", greeting)
	fmt.Printf("Number: %d\n", number)
}

func main() {

}

/**占位符
%v	值的默认格式表示。当输出结构体时，扩展标志（%+v）会添加字段名
%#v	值的Go语法表示
%T	值的类型的Go语法表示
%%	百分号
%t	单词true或false
%b	表示为二进制
%c	该值对应的unicode码值
%d	表示为十进制
%o	表示为八进制
%q	该值对应的单引号括起来的go语法字符字面值，必要时会采用安全的转义表示
%x	表示为十六进制，使用a-f
%X	表示为十六进制，使用A-F
%U	表示为Unicode格式：U+1234，等价于"U+%04X"
%b	无小数部分、二进制指数的科学计数法，如-123456p-78；参见strconv.FormatFloat
%e	科学计数法，如-1234.456e+78
%E	科学计数法，如-1234.456E+78
%f	有小数部分但无指数部分，如123.456
%F	等价于%f
%g	根据实际情况采用%e或%f格式（以获得更简洁、准确的输出）
%G	根据实际情况采用%E或%F格式（以获得更简洁、准确的输出）
%s	直接输出字符串或者[]byte
%q	该值对应的双引号括起来的Go语法字符串字面值，必要时会采用安全的转义表示
%x	每个字节用两字符十六进制数表示（使用a-f）
%X	每个字节用两字符十六进制数表示（使用A-F）
%p	表示为十六进制，并加上前导的0x
%f	默认宽度，默认精度
%9f	宽度9，默认精度
%.2f	默认宽度，精度2
%9.2f	宽度9，精度2
%9.f	宽度9，精度0
+	总是输出数值的正负号；对%q（%+q）会生成全部是ASCII字符的输出（通过转义）；
-	在输出右边填充空白而不是默认的左边（即从默认的右对齐切换为左对齐）；
#	切换格式：八进制数前加0（%#o），十六进制数前加0x（%#x）或0X（%#X），指针去掉前面的0x（%#p）； 对%q（%#q），如果strconv.CanBackquote返回真会输出反引号括起来的未转义字符串； 对%U（%#U），如果字符是可打印的，会在输出Unicode格式、空格、单引号括起来的Go字面值；
' '	对数值，正数前加空格而负数前加负号；对字符串采用%x或%X时（% x或% X）会给各打印的字节之间加空格；
0	使用0而不是空格填充，对于数值类型会把填充的0放在正负号后面；


*/
