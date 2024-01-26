package assignment

import "fmt"

const (
	//为了测试方便设置较小阈值
	minThresh = 8
	maxThresh = 64
)

func DeleteAt[T any](src []T, idx int) []T {
	length := len(src)
	if idx < 0 || idx >= length {
		panic("index out of range, can't delete element")
	} else {
		//从删除的目标下标开始，将元素往前挪动一位
		for i := idx; i < length-1; i++ {
			src[i] = src[i+1]
		}
		return checkAndShrink(src[:length-1])
	}
}

func checkAndShrink[T any](src []T) []T {
	c, shrinked := isShrinkRequired(src)
	if !shrinked {
		return src
	}
	ret := make([]T, 0, c)
	ret = append(ret, src...)
	fmt.Printf("shrinked %+v, len(%d), cap(%d) \n", ret, len(ret), cap(ret))
	return ret
}

func isShrinkRequired[T any](src []T) (int, bool) {
	var shrinked bool
	c, l := cap(src), len(src)
	var size = c
	if c >= minThresh {
		//容量超过最小容量阈值时才考虑缩容，容量小时对性能影响可以忽略，缩容反而可能造成额外性能损耗
		if c >= maxThresh && c/l >= 2 {
			//容量较大时，元素不足容量一半开始缩容
			size = c / 2
			shrinked = true
		} else if c < maxThresh && c/l >= 4 {
			//容量没超过最大阈值，元素不足1/4才缩容，避免过于频繁缩容
			size = c / 2
			shrinked = true
		}
	}
	return size, shrinked
}
