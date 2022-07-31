package intertest

import (
	"fmt"
	"testing"
)

//非入侵性，实现不依赖接口定义
//接口定义可以包含在使用包中
type Programmar interface {
	WriteHello() string
}

type GoProgrammar struct {
}

//DuckType
func (g *GoProgrammar) WriteHello() string {
	return ""
}

//空interface
func DoSomthing(p interface{}) {
	//断言
	if i, ok := p.(int); ok {
		fmt.Println("int:", i)
	}

	//switch
	switch s := p.(type) {
	case int:
		fmt.Println("swInt", s)
	default:
		fmt.Println("def:", s)
	}
}

func TestInter(t *testing.T) {
	// var p Programmar
	// p = new(GoProgrammar)
	// t.Log(p.WriteHello())
	DoSomthing(1)
}
