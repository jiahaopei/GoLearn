package functest

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

//可返回多返回值
func returnMutil() (int, int) {
	return rand.Intn(10), rand.Intn(20)
}

//可作为入参、返回值
func timeSpend(inner func(op int) int) func(op int) int {
	return func(op int) int {
		start := time.Now()
		ret := inner(op)
		fmt.Print("time:", time.Since(start).Seconds())
		return ret
	}
}

//可变长参数
func ssum(ops ...int) int {
	ret := 0
	for _, op := range ops {
		ret += op
	}
	return ret
}

//defer延迟执行

func TestFunc(t *testing.T) {
	defer func() {
		t.Log("clear")
	}()
	a, _ := returnMutil()
	t.Log(a)
}
