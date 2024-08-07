package contas

import (
	"curso-go-poo/pkg/clientes"
	"fmt"
)

type ContaPoupanca struct {
	Titular 								clientes.Titular
	NumeroAgencia, NumeroConta, Operacao 	int
	saldo 									float64
}


func (c *ContaPoupanca) Sacar(valorDoSaque float64) string {

	podeSacar := valorDoSaque <= c.saldo && valorDoSaque > 0

	if podeSacar {
		c.saldo -= valorDoSaque
		return "Saque realizado com sucesso!"
	} else {
		return "saldo insuficiente."
	}
}

func (c *ContaPoupanca) Depositar(valorDoDeposito float64) (string, float64) {
	podeDepositar := valorDoDeposito > 0

	if podeDepositar {
		c.saldo += valorDoDeposito
		return "Deposito realizado com sucesso!", c.saldo
	} else {
		return "Não foi possível realizar o deposito. Entre com um valor válido.", c.saldo
	}
}

func (c *ContaPoupanca) Transferir(valorDaTransferencia float64, contaDestino *ContaPoupanca) string {
	podeTransferir := valorDaTransferencia > 0 && c.saldo > valorDaTransferencia

	if podeTransferir {
		c.saldo -= valorDaTransferencia
		contaDestino.Depositar(valorDaTransferencia)
		return "Transferencia realizada com sucesso!"
	} else {
		return "Ocorreu um erro na transferência. Por gentileza, cheque se os campos inseridos estão corretos"
	}
}

func (c *ContaPoupanca) ConsultarSaldo() string {
	return "Seu saldo é " + fmt.Sprintf("%.2f", c.saldo)
}
