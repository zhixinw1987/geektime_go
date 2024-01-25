package functions

func CallNames(name string, alias ...string) {
	println("hello, my name is " + name)
	if len(alias) > 0 {
		for _, v := range alias {
			println("aka " + v)
		}
	}
}
