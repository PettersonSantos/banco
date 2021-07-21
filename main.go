package main

import "fmt"

type ContaCorrente struct {
	titular string
	numeroAgencia int
	numeroConta int
	saldo float32
}

func (c *ContaCorrente) Sacar(valorDoSaque float32) string {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.saldo
	if !podeSacar {
		return "Saldo Insuficiente"
	}
	c.saldo -= valorDoSaque
	return "Saldo Realizado com Sucesso"
}

func (c *ContaCorrente) DepositarSaldo(valorDoDeposito float32) (string, float32) {
	if valorDoDeposito > 0 {
		c.saldo += valorDoDeposito
		return "Depósito Realizado com Sucesso", c.saldo
	}

	return "O valor do depósito menor que 0", c.saldo
}

func (c *ContaCorrente) Transferir(valorTransferencia float32, contaDestino *ContaCorrente) bool {
	if valorTransferencia < c.saldo && valorTransferencia > 0{
		c.saldo -= valorTransferencia
		contaDestino.DepositarSaldo(valorTransferencia)
		return true
	}
	return false
}

func main() {
	correntePetterson := ContaCorrente{
		titular:       "Petterson",
		numeroAgencia: 556,
		numeroConta:   56578,
		saldo:         500,
	}

	correnteNew := ContaCorrente{
		titular:       "João",
		numeroAgencia: 556,
		numeroConta:   56578,
		saldo:         300,
	}

	status := correnteNew.Transferir(200, &correntePetterson)

	fmt.Println(status)
	fmt.Println(correntePetterson.saldo)
	fmt.Println(correnteNew.saldo)
}
