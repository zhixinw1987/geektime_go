package types

type GenericList[T any] interface {
	Add(idx int, t T)
	Append(t T)
}

func UseGenericList() {
	var list GenericList[int]
	list.Append(1)
	list.Append(2)
	// list.Append(2.2)
}

type GLinkedList[T any] struct {
	head *node[T]
	e    T
}

type node[T any] struct {
	val T
}

func SumNumber[T Number](data ...T) T {
	var res T
	for _, v := range data {
		res = res + v
	}
	return res
}

type Number interface {
	int | uint | int8 | uint8 | int16 | uint16
}
