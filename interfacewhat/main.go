// Package main doing main things
package main

func main() {
	var x *string
	x = nil
	isNull(x)
}

func isNull(i interface{}) {
	if i == nil {
		println("is null")
	} else {
		println("not null")
	}
}

func isNulls(is ...interface{}) {
	for idx := range is {
		if is[idx] == nil {
			println("is null")
		} else {
			println("not null")
		}
	}
}
