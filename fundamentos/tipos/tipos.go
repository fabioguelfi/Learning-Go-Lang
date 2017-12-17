package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	// numericos inteiros
	fmt.Println(1, 2, 1000)

	fmt.Println("literal inteiro e: ", reflect.TypeOf(32000))

	// sem sinal(so valores positivos ...) uint8 uint16 uint32 uint64

	var b byte = 255

	fmt.Println("o byte e", reflect.TypeOf(b))

	// com sinal... int8 int16 int32 int64

	i1 := math.MaxInt64
	i2 := math.MaxInt32
	i3 := math.MaxInt16
	i4 := math.MaxInt8

	fmt.Println("O valor maximo do int64 e:", i1)
	fmt.Println("E o valor do tipo da variavel i1 e:", reflect.TypeOf(i1))

	fmt.Println("O valor maximo do int32 e:", i2)
	fmt.Println("O valor maximo do int16 e:", i3)
	fmt.Println("O valor maximo do int8 e", i4)

	var i5 rune = 'a' // representa um mapeamento da tabela Unicode (int32)

	fmt.Println("O rune e:", reflect.TypeOf(i5))
	fmt.Println(i5)

	// numeros reais (float32, float64)
	var x float32 = 49.99
	var y float64 = 49.99
	fmt.Println("O tipo da variavel x e:", reflect.TypeOf(x))
	fmt.Println("O tipo literal default float e:", reflect.TypeOf(49.99)) //float64
	fmt.Println("O tipo da variavel y e:", reflect.TypeOf(y))

	// boolean

	bo := true
	fmt.Println("O tipo da variavel bo:", reflect.TypeOf(bo))
	fmt.Println("O valor da variavel bo e:", bo)

	// string
	s1 := "String de testes"
	fmt.Println(s1, "!")
	fmt.Println("O tamanho da string e", len(s1))

	// string com multiplas linhas
	s2 := `
	ola
	meu
	nome
	e
	fabio
	guelfi
	`

	fmt.Println(s2, len(s2), reflect.TypeOf(s2))

	// char ?

	char := 'a'
	fmt.Println("O tipo de char e:", reflect.TypeOf(char), "nao existe o tipo char")

}
