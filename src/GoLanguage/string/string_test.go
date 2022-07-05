package stringtest

import (
	"strconv"
	"strings"
	"testing"
)

//只读、len为byte长度
func TestString(t *testing.T) {
	var s string
	s = "hello"
	t.Log(len(s))

	s = "\xE4\xB8\xFF"
	t.Log(s, len(s))

	s = "中"
	c := []rune(s)
	t.Logf("unicode:%x", c[0])
	t.Logf("UTF-8:%x", s)

	s = "中国"
	for _, c := range s {
		t.Logf("%[1]c,%[1]d", c)
	}

}

func TestStrFunc(t *testing.T) {
	s := "a,b,c"
	parts := strings.Split(s, ",")
	for _, c := range parts {
		t.Log(c)
	}
	t.Log(strings.Join(parts, "-"))

	e := strconv.Itoa(10)
	t.Log(e)
	if i, err := strconv.Atoi("10"); err == nil {
		t.Log(i)
	}
}
