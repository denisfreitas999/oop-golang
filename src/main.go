package main

import (
	"curso-go-poo/pkg/clientes"
	"curso-go-poo/pkg/contas"
	"fmt"
)

func PagarBoleto(conta contas.Conta, valorDoBoleto float64) {
	conta.Sacar(valorDoBoleto)
}

func main() {
	fmt.Println("=========== Conta da Cris ===========")
	// Cliente Cris
	clienteCris := &clientes.Titular{
		CPF:       "063.580.380-10",
		Nome:      "Cris Souza",
		Profissao: "Médica",
	}
	// Conta da Cris
	contaDaCris := &contas.ContaCorrente{
		ContaBase: contas.ContaBase{
			Titular:       clienteCris,
			NumeroAgencia: 2654,
			NumeroConta:   12456,
		},
	}
	contaDaCris.Depositar(12503.29)
	fmt.Println(*contaDaCris.Titular)
	fmt.Println(*contaDaCris)
	fmt.Println(contaDaCris.ConsultarSaldo())

	fmt.Println("=========== Conta do Denisson ===========")
	// Cliente Denisson
	clienteDenisson := &clientes.Titular{
		CPF:       "538.240.490-90",
		Nome:      "Denisson Freitas",
		Profissao: "Engenheiro DevOps",
	}
	// Conta do Denisson
	contaDoDenisson := &contas.ContaPoupanca{
		ContaBase: contas.ContaBase{
			Titular:       clienteDenisson,
			NumeroAgencia: 1222,
			NumeroConta:   35694,
		},
		Operacao: 5,
	}
	contaDoDenisson.Depositar(43520.25)
	fmt.Println(*contaDoDenisson.Titular)
	fmt.Println(*contaDoDenisson)
	fmt.Println(contaDoDenisson.ConsultarSaldo())

	fmt.Println("=========== Conta do Denisson : Pagar Boleto ===========")
	PagarBoleto(contaDoDenisson, 520.25)
	fmt.Println(contaDoDenisson.ConsultarSaldo())
}
