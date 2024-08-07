package contas

type ContaCorrente struct {
	Titular       string
	NumeroAgencia int
	NumeroConta   int
	Saldo         float64
}

func (c *ContaCorrente) Sacar(valorDoSaque float64) string {

	podeSacar := valorDoSaque <= c.Saldo && valorDoSaque > 0

	if podeSacar {
		c.Saldo -= valorDoSaque
		return "Saque realizado com sucesso!"
	} else {
		return "Saldo insuficiente."
	}
}

func (c *ContaCorrente) Depositar(valorDoDeposito float64) (string, float64) {
	podeDepositar := valorDoDeposito > 0

	if podeDepositar {
		c.Saldo += valorDoDeposito
		return "Deposito realizado com sucesso!", c.Saldo
	} else {
		return "Não foi possível realizar o deposito. Entre com um valor válido.", c.Saldo
	}
}

func (c *ContaCorrente) Transferir(valorDaTransferencia float64, contaDestino *ContaCorrente) string {
	podeTransferir := valorDaTransferencia > 0 && c.Saldo > valorDaTransferencia

	if podeTransferir {
		c.Saldo -= valorDaTransferencia
		contaDestino.Depositar(valorDaTransferencia)
		return "Transferencia realizada com sucesso!"
	} else {
		return "Ocorreu um erro na transferência. Por gentileza, cheque se os campos inseridos estão corretos"
	}
}