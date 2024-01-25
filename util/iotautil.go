package util

const (
	index0 = iota
	index1
	index2
	index3
)

const (
	Monday = iota + 1
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday
)

func ShowIndex() {
	println(index0, index1, index2)
}

func ShowWeekday() {
	println(Monday, Tuesday, Wednesday)
}
