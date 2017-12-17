package main

import (
	"fmt"
	conv "strconv"
)

func main() {
	x := 2.4
	y := 2

	fmt.Println(x / float64(y))

	nota := 6.9
	notaFinal := int(nota)
	fmt.Println(notaFinal)

	//cuidado
	fmt.Println("Teste" + string(97)) // 97 e equivalente a na tabela asc

	//int para string
	fmt.Println("Teste:", conv.Itoa(123))

	//string para int
	num, _ := conv.Atoi("7")

	fmt.Println(num - 6)

	b, _ := conv.ParseBool("true")

	if b {
		fmt.Println("verdadeiro")
	} else {
		fmt.Println("falso")
	}
}
