package series

import (
	"errors"
	"fmt"
)

//在main方法执行前执行，不同包的init方法按照依赖关系决定执行顺序，可以有多个init方法
func init() {
	fmt.Println("")
}

//大写可访问,小写则不可访问fib
func Fib(n int) ([]int, error) {
	if n < 2 || n > 100 {
		return nil, errors.New("dfd")
	}
	fibList := []int{1, 1}
	for i := 2; i < n; i++ {
		fibList = append(fibList, fibList[i-1]+fibList[i-2])
	}
	return fibList, nil
}
