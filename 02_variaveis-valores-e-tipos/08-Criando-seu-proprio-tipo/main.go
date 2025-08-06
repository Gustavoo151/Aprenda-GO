package main

import "fmt"

type hotdog int

var b hotdog

func main() {
	x := 10
	fmt.Printf("%T\n", x) // Imprime o tipo de x, que é int
	fmt.Printf("%T\n", b) // Imprime o valor zero de hotdog, que é 0

	//x = b // Isso não é permitido, pois hotdog e int são tipos diferentes
}
