package main

import (
	"fmt"
)

func main() {
	//numerosDeBytes, erros := fmt.Println("Hello World")
	numerosDeBytes, _ := fmt.Println("Hello World") // A função Println retorna dois valores e somo obrigados a pegalos
	// mas se não usarmos os dois valores o código não compila. Um solução para isso é usar o _ (Black identifier) que sinaliza
	// que aquilo é "lixo"

	//fmt.Println("Número de Bytes: ", numerosDeBytes, "Erros: ", erros)
	fmt.Println("Número de Bytes: ", numerosDeBytes)

}
