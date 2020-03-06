package to

import (
	"fmt"
	"testing"
)

type Student struct {
	Age   int
	Name  string
	Class string
}

func TestString(t *testing.T) {
	str := String(-1)
	fmt.Println(str)

	v := &Student{
		Age:   12,
		Name:  "Jason Bourne",
		Class: "CIA",
	}

	fmt.Println(HexString(v))
}
