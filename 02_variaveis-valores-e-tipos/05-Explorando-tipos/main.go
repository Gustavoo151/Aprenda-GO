package main

import "fmt"

// Go é uma linguagem de tipagem estatica

var x int = 10 // Aqui estamos declarando o tipo
var y = 10.5   // Aqui o tipo de é difinido automaticamente

func main() {
	//x = 20.5 // Isso dá erro, pois estamos colocando o tipo float em uma variável float
	fmt.Printf("%v, %T\n", x, x)
	fmt.Printf("%v, %T\n", y, y)

}
