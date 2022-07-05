package slicetest

import "testing"

//切片可伸缩，不可比较。数组不可伸缩，可比较
//[]int{}             [...]int{}
func TestSlice(t *testing.T) {
	var s0 []int
	t.Log(len(s0), cap(s0))
	s0 = append(s0, 1)
	t.Log(len(s0), cap(s0))

	s1 := []int{1, 2, 3}
	t.Log(len(s1), cap(s1))

	s2 := make([]int, 3, 5)
	t.Log(len(s2), cap(s2))
}

func TestSliceShare(t *testing.T) {
	year := []string{"a", "b", "c"}
	q1 := year[1:2]
	t.Log(q1)
	t.Log(len(q1), cap(q1))
	q1[0] = "d"
	t.Log(year)
}
