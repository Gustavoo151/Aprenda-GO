package main

import "fmt"

var x int

func main() {
	x = 10 // isso é uma inicialização, é a primeira atribuição de valor em variavl
	fmt.Printf("%v, %T\n", x, x)
	x = 20 // esse valor é uma atribuição
	fmt.Printf("%v, %T\n", x, x)
	//x := 20 // O operador curto de atribuição já faz essas duas parte a inicialização e atribuição

	/*ESSES são ps valores defaul de cada tipo
		- int: 0
		- float: 0.0
		- boolean: false
	 	- string: ""
		- pointers, functions, interfaces, slices, channels, maps: nil
	*/

	// Use sempre o operado curto de atribuição :=
	// User var quando for usar alguma variável a nivel global
}
