package functions

func func1() {

}

//go 不支持重载
// func func1 (string a) {

// }

// 同类型参数可以连续申明
// 支持多个返回值
func func2(a, b, c string) (ret1, ret2 string) {
	if len(a) >= len(b) {
		ret1 = a
	} else {
		ret1 = c
	}
	ret2 = a + b + c
	return ret1, ret2
}

// 函数式编程，将方法当作变量调用
func wrappedPrinter(args string) {
	println("this is message: ", args)
}

func Print(message string) {
	myFunc := wrappedPrinter
	myFunc(message)

}

// 匿名方法
func AnonymousFunc(message string) {
	//将方法赋值给一个变量，通过变量调用方法
	fn := func(content string) {
		println("this is message from anonymous func: ", content)
	}
	fn(message)

	//定义匿名方法后直接发起调用
	ret := func(string) string {
		return "anonymous func " + message
	}(message)
	println(ret)
}
