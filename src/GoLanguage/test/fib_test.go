package try_test

import (
	"fmt"
	"testing"
)

func TestFib(t *testing.T) {
	// var (
	// 	a int = 1
	// 	b int = 1
	// )

	var (
		a = 1
		b = 1
	)

	// a:=1
	// b:=1
	// t.Log(a)
	fmt.Print(a)
	for i := 0; i < 5; i++ {
		fmt.Print("", b)
		tmp := a
		a = b
		b = tmp + a
	}
	fmt.Println()
}
