package types

import "fmt"

type List interface {
	Add(idx int, val any) error
	Del(idx int)
	Find(val any) any
}

type ArrayList struct {
	elements []any
}

func (l *ArrayList) Add(idx int, val any) error {
	if idx > len(l.elements) {
		panic("idx out of range")
	} else {
		l.elements[idx] = val
	}
	return nil
}

func (l *ArrayList) Del(idx int) {
	if idx > len(l.elements) {
		panic("idx out of range")
	} else {
		temp := make([]any, len(l.elements)-1)
		for i := 0; i < len(l.elements); i++ {
			if i == idx {
				continue
			}
			temp[i] = l.elements[i]
		}
		l.elements = temp
	}
}

func (l *ArrayList) Find(val any) any {
	for _, v := range l.elements {
		if v == val {
			return val
		}
	}
	return nil
}

func UseList(list List) {
	fmt.Printf(`Go没有implement关键字，一个struct实现了接口的所有方法，这个struct自然就实现了该接口`)
	list.Find("abc")
}
