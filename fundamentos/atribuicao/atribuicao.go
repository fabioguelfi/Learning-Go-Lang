package main

import (
	"fmt"
	"reflect"
)

func main() {
	var b byte = 3
	fmt.Println(b)

	i := 3 //inferencia de tipo automatica
	i += 3 // i = i + 3
	i -= 3 // i = i - 3
	i /= 3 // i = i / 2
	i *= 2 // i = i * 2
	i %= 2 // i = i % 2
	fmt.Println(i)

	x, y := 1, 2
	fmt.Println(x, y)

	x, y = y, x
	fmt.Println(reflect.TypeOf(x), x, reflect.TypeOf(y), y)
}
