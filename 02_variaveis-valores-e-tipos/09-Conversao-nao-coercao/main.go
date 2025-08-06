package main

import "fmt"

type hotdog int

var b hotdog = 101

func main() {
	x := 10
	fmt.Printf("%T\n", x) // Imprime o tipo de x, que é int
	fmt.Printf("%T\n", b) // Imprime o valor zero de hotdog, que é 0

	x = int(b) // Aqui estamos convertendo b de hotdog para int, não existe casting em Go, mas sim conversão de tipos
}
