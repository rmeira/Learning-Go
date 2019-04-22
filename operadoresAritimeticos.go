package main

import (
	"fmt"
)

func main(){
	var numero1 int
	var numero2 int
	fmt.Print("Digite o primeiro numero: ")
	// o & comercial fala que o valor deve ser armazenado na mesma posição da memoria
	fmt.Scanln(&numero1)

	fmt.Print("Digite o segundo numero: ")
	fmt.Scanln(&numero2)

	fmt.Printf("Adição dos numeros: %d + %d = %d \n", numero1, numero2, numero1+numero2)
	fmt.Printf("Subtração dos numeros: %d - %d = %d \n", numero1, numero2, numero1-numero2)
	fmt.Printf("Multiplicação dos numeros: %d * %d = %d \n", numero1, numero2, numero1*numero2)
	fmt.Printf("Divisão dos numeros: %d / %d = %d \n", numero1, numero2, numero1/numero2)
	fmt.Printf("Tipo dos numeros: %d %% %d = %d \n", numero1, numero2, numero1%numero2)
}