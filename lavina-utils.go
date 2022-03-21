package lvn

import "fmt"

func Default(s interface{}) {
	fmt.Println("")
}

func ternary(c bool, t interface{}, f interface{}) interface{} {
	if c {
		return t
	} else {
		return f
	}
}
