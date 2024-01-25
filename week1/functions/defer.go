package functions

//Defer is like finally in Java, execute the code block before exiting
//Defer define first execute last
func Defer() {
	println("open DS")
	defer func() {
		println("defer1: close DS")
	}()
	defer func() {
		println("defer2: close tranaction")
	}()
	println("start DB query")
	println("start transaction")
}
