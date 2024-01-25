package main

import (
	"fmt"
	"unicode/utf8"

	"github.com/zhixinw1987/geektime_go/util"
	"github.com/zhixinw1987/geektime_go/week1/collections"
	"github.com/zhixinw1987/geektime_go/week1/control"
	"github.com/zhixinw1987/geektime_go/week1/functions"
	"github.com/zhixinw1987/geektime_go/week1/types"
)

func main() {
	var str1 string = "test string"
	var str2 string = "测试"
	var num1 uint = 256
	print("concat string with different data type: ")
	println(str1 + fmt.Sprint(num1))
	print("concat string with string: ")
	println(str1 + str2)
	print("count string: ")
	println(len(str1) + len(str2))
	println(len(str1) + utf8.RuneCountInString(str2))
	println(util.Global)

	util.ShowIndex()
	util.ShowWeekday()
	functionsDemo()
	functions.Defer()
	controlDemo()
	collDemo()
	functions.DeferLoopV1()
	functions.DeferLoopV2()
	functions.DeferLoopV3()
	typeDemo()
}

func functionsDemo() {
	println("functions demo===================")
	functions.Print("main")
	functions.AnonymousFunc("main")

	println("closure demo==================")
	count := functions.Closure()
	println(count())
	println(count())
	println(count())

	count = functions.Closure()
	println(count())
	println(count())
	println(count())

	println("variable params=================")
	functions.CallNames("demo", "a", "abs", "ttt")

	alias := []string{"b", "bb", "bbb"}
	functions.CallNames("B", alias...)
}

func controlDemo() {
	println("control demo====================")
	control.ForLoop()
	control.WhileLoop()
	control.ForLoop()
	control.MapLoop()
	control.SwitchVal(10)
	control.SwitchExpr(150)
}

func collDemo() {
	println("collections demo======================")
	collections.Array()
	collections.Slice()
	collections.SubSlice()
}

func typeDemo() {
	types.NewUser()
	types.ChangeUser()
}
