package maptest

import "testing"

func TestMap(t *testing.T) {
	map1 := map[int]int{1: 1, 2: 0}
	map2 := map[int]int{}
	map3 := make(map[int]int, 10)
	t.Logf("len m1=%d,m2=%d,m3=%d", len(map1), len(map2), len(map3))

	if v, ok := map1[2]; ok {
		t.Log(v + map1[2])
	} else {
		t.Logf("nothing:%d", v)
	}

	for i, v := range map1 {
		t.Log(i, v)
	}
}

func TestMapExt(t *testing.T) {
	m := map[int]func(op int) int{}
	m[1] = func(op int) int { return op }
	t.Log(m[1](2))
}

//map实现set
func TestMapSet(t *testing.T) {
	m := map[int]bool{1: true}
	delete(m, 1)
}
