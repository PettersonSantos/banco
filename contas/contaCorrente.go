package contas

import "banco/clientes"

type ContaCorrente struct {
	Titular                    clientes.Titular
	NumeroAgencia, NumeroConta int
	saldo                      float32
}

func (c *ContaCorrente) Sacar(valorDoSaque float32) string {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.saldo
	if !podeSacar {
		return "saldo Insuficiente"
	}
	c.saldo -= valorDoSaque
	return "saldo Realizado com Sucesso"
}

func (c *ContaCorrente) DepositarSaldo(valorDoDeposito float32) (string, float32) {
	if valorDoDeposito > 0 {
		c.saldo += valorDoDeposito
		return "Depósito Realizado com Sucesso", c.saldo
	}

	return "O valor do depósito menor que 0", c.saldo
}

func (c *ContaCorrente) Transferir(valorTransferencia float32, contaDestino *ContaPoupanca) bool {
	if valorTransferencia < c.saldo && valorTransferencia > 0 {
		c.saldo -= valorTransferencia
		contaDestino.DepositarSaldo(valorTransferencia)
		return true
	}
	return false
}

func (c *ContaCorrente) ObterSaldo() float32 {
	return c.saldo
}
