package main

import (
	"fmt"
	"banco/contas"
)



func main() {
	correntePetterson := contas.ContaCorrente{
		Titular:       "Petterson",
		NumeroAgencia: 556,
		NumeroConta:   56578,
		Saldo:         500,
	}

	correnteNew := contas.ContaCorrente{
		Titular:       "João",
		NumeroAgencia: 556,
		NumeroConta:   56578,
		Saldo:         300,
	}

	status := correnteNew.Transferir(200, &correntePetterson)

	fmt.Println(status)
	fmt.Println(correntePetterson.Saldo)
	fmt.Println(correnteNew.Saldo)
}
