package encaptest

import (
	"fmt"
	"testing"
)

type Emp struct {
	Name string
	Id   string
	Age  int
}

//实例对应方法被调用时，会进行值复制
func (e Emp) String() string {
	return fmt.Sprintf("Id:%s,Name:%s,Age:%d", e.Id, e.Name, e.Age)
}

//使用指针避免内存拷贝
func (e *Emp) StringF() string {
	return fmt.Sprintf("Id:%s,Name:%s,Age:%d", e.Id, e.Name, e.Age)
}

func TestEncap(t *testing.T) {
	e := Emp{"aab", "123", 21}
	t.Log(e.String())
}
