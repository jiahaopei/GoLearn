package err_test

import (
	"errors"
	"testing"
)

var RangeError = errors.New("n in [2,100]")

func Fib(n int) ([]int, error) {
	if n < 2 || n > 100 {
		return nil, RangeError
	}
	fibList := []int{1, 1}
	for i := 2; i < n; i++ {
		fibList = append(fibList, fibList[i-1]+fibList[i-2])
	}
	return fibList, nil
}

func TestErr(t *testing.T) {
	if v, err := Fib(3); err != nil {
		if err == RangeError {
			t.Error("error:", err)
		}

	}
}
