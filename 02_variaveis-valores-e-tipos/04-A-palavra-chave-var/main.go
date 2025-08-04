package main

import "fmt"

var z = 10 //  Não podemos ultilizar variáveis com escopo global usando o := , com isso temos  que usar
//o var para criar a variável

func main() {
	y := 10
	qualquerCoisa(y)
}

func qualquerCoisa(x int) {
	fmt.Println(x)
}
