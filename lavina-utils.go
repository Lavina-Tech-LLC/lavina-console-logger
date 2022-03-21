package lvn

import "fmt"

func Default(s interface{}) {
	fmt.Println("")
}

func Ternary(c bool, t interface{}, f interface{}) interface{} {
	if c {
		return t
	} else {
		return f
	}
}
