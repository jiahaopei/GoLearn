package goroutine_test

import (
	"fmt"
	"testing"
	"time"
)

func TestGoRoutine(t *testing.T) {
	for i := 0; i < 5; i++ {
		go func(i int) {
			fmt.Println(i)
		}(i)
	}
	time.Sleep(time.Millisecond * 40)
}
