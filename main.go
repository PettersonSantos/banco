package main

import "fmt"

type ContaCorrente struct {
	titular string
	numeroAgencia int
	numeroConta int
	saldo float32
}

func main() {
	corrente := ContaCorrente{
		titular:       "Petterson",
		numeroAgencia: 556,
		numeroConta:   56578,
		saldo:         535.2,
	}

	fmt.Println(corrente)
}
