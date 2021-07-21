package contas

import "banco/clientes"

type ContaPoupanca struct {
	Titular                              clientes.Titular
	NumeroAgencia, NumeroConta, Operacao int
	saldo                                float32
}

func (c *ContaPoupanca) Sacar(valorDoSaque float32) string {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.saldo
	if !podeSacar {
		return "saldo Insuficiente"
	}
	c.saldo -= valorDoSaque
	return "saldo Realizado com Sucesso"
}

func (c *ContaPoupanca) DepositarSaldo(valorDoDeposito float32) (string, float32) {
	if valorDoDeposito > 0 {
		c.saldo += valorDoDeposito
		return "Depósito Realizado com Sucesso", c.saldo
	}

	return "O valor do depósito menor que 0", c.saldo
}

func (c *ContaPoupanca) ObterSaldo() float32 {
	return c.saldo
}
