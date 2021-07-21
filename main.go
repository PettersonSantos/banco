package main

import (
	"banco/contas"
	"fmt"
)

func main() {
	//clientePetterson := clientes.Titular{Nome: "Petterson", CPF: "12345678901", Profissao: "Dev"}
	contaDoPetterson := contas.ContaCorrente{}
	contaDoPetterson.DepositarSaldo(100)
	PagarBoleto(&contaDoPetterson, 40)
	fmt.Println(contaDoPetterson.ObterSaldo())

}

func PagarBoleto(conta verificarConta, valorDoBoleto float32) {
	conta.Sacar(valorDoBoleto)
}

type verificarConta interface {
	Sacar(valor float32) string
}
