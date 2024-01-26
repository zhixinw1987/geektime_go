package assignment_test

import (
	"testing"

	"github.com/zhixinw1987/geektime_go/week1/assignment"
)

func TestShrink(t *testing.T) {
	var sli []int
	for i := 0; i < 200; i++ {
		sli = append(sli, i)
	}

	for i := 0; i < 150; i++ {
		sli = assignment.DeleteAt(sli, 3)
	}
}
