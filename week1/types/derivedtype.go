package types

type Integer int

func ConvertInt() {
	i1 := 10
	i2 := Integer(i1)
	i3 := int(i2)
	println(i3)
}

type Fish struct{ Weight float32 }

func (f *Fish) Swim() {
	println("swiming ...")
}

type FakeFish Fish

func UseFish() {
	f1 := Fish{1.2}
	f2 := FakeFish{}
	f1.Swim()
	println(f2.Weight)
	//f2.Swim()
	println("衍生类型和原类型可以互相转换，衍生类型可以访问原类型属性，但是不能调用原类型定义的接收器")

	println("定义类型的别名，和原类型完全一样")
	f3 := Yu{Weight: 2}
	f3.Swim()
}

type Yu = Fish
