package main

import (
	"banco/contas"
	"fmt"
)

func main() {
	//clientePetterson := clientes.Titular{Nome: "Petterson", CPF: "12345678901", Profissao: "Dev"}
	contaDoPetterson := contas.ContaCorrente{}
	contaDoPetterson.DepositarSaldo(100)
	fmt.Println(contaDoPetterson.ObterSaldo())


}
