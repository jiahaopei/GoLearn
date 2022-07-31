package csp_test

import (
	"fmt"
	"testing"
	"time"
)

func service() string {
	return "Done"
}

func otherService() {

}

func AsyncService() chan string {
	retCh := make(chan string)

	go func() {
		ret := service()
		retCh <- ret
	}()
	return retCh
}

func TestCsp(t *testing.T) {
	retCh := AsyncService()
	otherService()
	fmt.Println(<-retCh)
}

//多路选择
func TestSelect(t *testing.T) {
	select {
	case ret := <-AsyncService():
		t.Log(ret)
	case <-time.After(time.Millisecond * 100):
		t.Error("time out")
	}
}
