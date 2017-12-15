package main

import (
	"fmt"
)

func main() {
	fmt.Print("Mesma ")
	fmt.Print("Linha\n")

	fmt.Println("Nova")
	fmt.Println("Linha")

	x := 3.141333

	//fmt.Println("O valor de x é de", x)

	xs := fmt.Sprint(x)
	fmt.Println("O valor de x é " + xs)
	fmt.Println("O valor de x é ", x)

	name := "Fabio Guelfi"

	fmt.Printf("O valor de x é %.2f %s .", x, name)

	a := 1
	b := 1.999
	c := false
	d := "opa"
	fmt.Printf("\n %d %f %.1f %t %s", a, b, b, c, d)

}
