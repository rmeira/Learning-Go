package main

import (
	"fmt"
)

func main() {
	fmt.Println("Declarando numeros inteiros e sem sinal")

	var numberUint8 uint8 = 255
	fmt.Println(numberUint8)

	var numberByte byte = 255
	fmt.Println(numberByte)

	var numberUint16 uint16 = 10000
	fmt.Println(numberUint16)

	var numberUint32 uint32 = 100001231
	fmt.Println(numberUint32)

	var numberUint64 uint64 = 10000123123123333333
	fmt.Println(numberUint64)


	fmt.Println("Declarando numeros inteiros com sinal")

	var numberInt8 int8 = -128
	fmt.Println(numberInt8)

	var numberInt16 int16 = -128
	fmt.Println(numberInt16)

	var numberInt32 int32 = -128
	fmt.Println(numberInt32)

	var numberRune rune = -128
	fmt.Println(numberRune)

	var numberInt64 int64 = -128
	fmt.Println(numberInt64)

	fmt.Println("Declarando numeros que assumem a arquitetura do processador")
	// Assume a arquitetura do processador, se a maquina rodas em 32, será 32 se for 64 será 64
	var numberUint uint = 128 // Apenas numeros positivos
	fmt.Println(numberUint)
	var numberint int = 128
	fmt.Println(numberint)

	fmt.Println("Declarando numeros de pontos flutuantes")
	// Assume a arquitetura do processador, se a maquina rodas em 32, será 32 se for 64 será 64
	var numberFloat32 float32 = 128.123
	fmt.Println(numberFloat32)
	var numberFloat64 float64 = 128.123
	fmt.Println(numberFloat64)
	var numberComplex64 complex64 = 128.123
	fmt.Println(numberComplex64)

	fmt.Println("Declarando Strings")
	var stringOne string = "Rafael"
	fmt.Println(stringOne)

	var stringTwo string =`Rafael
Meira`
	fmt.Println(stringTwo)

	fmt.Println("Declarando Boleanos")
	var boolean bool = true
	fmt.Println(boolean)

	fmt.Println("Multiplas declarações")
	var (
		nome 		string = "Rafael"
		sobrenome 	string = "Meira"
	)
	fmt.Println(nome + " " + sobrenome)
}
