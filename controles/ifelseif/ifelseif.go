package main

import "fmt"

func notaParaConceito(n float64) string {
	if n >= 9 && n <= 10 {
		return "A"
	} else if n >= 8 && n < 9 {
		return "B"
	} else if n >= 5 && n < 8 {
		return "C"
	} else if n >= 3 && n < 5 {
		return "D"
	} else if n >= 0 && n < 3 {
		return "E"
	}
	return "Dados Invalidos"
}

func main() {
	conceito := notaParaConceito(7.3)
	fmt.Println(conceito)
	conceito1 := notaParaConceito(5.3)
	fmt.Println(conceito1)
	conceito2 := notaParaConceito(0.1)
	fmt.Println(conceito2)
	conceito3 := notaParaConceito(9)
	fmt.Println(conceito3)
}
