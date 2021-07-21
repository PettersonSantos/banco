package contas

type ContaCorrente struct {
	Titular string
	NumeroAgencia int
	NumeroConta int
	Saldo float32
}

func (c *ContaCorrente) Sacar(valorDoSaque float32) string {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.Saldo
	if !podeSacar {
		return "Saldo Insuficiente"
	}
	c.Saldo -= valorDoSaque
	return "Saldo Realizado com Sucesso"
}

func (c *ContaCorrente) DepositarSaldo(valorDoDeposito float32) (string, float32) {
	if valorDoDeposito > 0 {
		c.Saldo += valorDoDeposito
		return "Depósito Realizado com Sucesso", c.Saldo
	}

	return "O valor do depósito menor que 0", c.Saldo
}

func (c *ContaCorrente) Transferir(valorTransferencia float32, contaDestino *ContaCorrente) bool {
	if valorTransferencia < c.Saldo && valorTransferencia > 0{
		c.Saldo -= valorTransferencia
		contaDestino.DepositarSaldo(valorTransferencia)
		return true
	}
	return false
}
