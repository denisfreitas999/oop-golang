package contas

import (
	"curso-go-poo/pkg/clientes"
	"fmt"
)

type ContaBase struct {
	Titular       *clientes.Titular
	NumeroAgencia int
	NumeroConta   int
	saldo         float64
}

func (c *ContaBase) Sacar(valorDoSaque float64) string {
	podeSacar := valorDoSaque <= c.saldo && valorDoSaque > 0
	if podeSacar {
		c.saldo -= valorDoSaque
		return "Saque realizado com sucesso!"
	} else {
		return "Saldo insuficiente."
	}
}

func (c *ContaBase) Depositar(valorDoDeposito float64) (string, float64) {
	podeDepositar := valorDoDeposito > 0
	if podeDepositar {
		c.saldo += valorDoDeposito
		return "Depósito realizado com sucesso!", c.saldo
	} else {
		return "Não foi possível realizar o depósito. Entre com um valor válido.", c.saldo
	}
}

func (c *ContaBase) Transferir(valorDaTransferencia float64, contaDestino Conta) string {
	podeTransferir := valorDaTransferencia > 0 && c.saldo > valorDaTransferencia
	if podeTransferir {
		c.saldo -= valorDaTransferencia
		contaDestino.Depositar(valorDaTransferencia)
		return "Transferência realizada com sucesso!"
	} else {
		return "Ocorreu um erro na transferência. Verifique se os campos inseridos estão corretos."
	}
}

func (c *ContaBase) ConsultarSaldo() string {
	return "Seu saldo é " + fmt.Sprintf("%.2f", c.saldo)
}

func (c ContaBase) String() string {
    return fmt.Sprintf("{%s %d %d %.2f}", c.Titular, c.NumeroAgencia, c.NumeroConta, c.saldo)
}