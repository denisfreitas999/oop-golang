package main

import (
	"curso-go-poo/pkg/clientes"
	"curso-go-poo/pkg/contas"
	"fmt"
)

func PagarBoleto(conta verificaConta, valorDoBoleto float64) {
	conta.Sacar(valorDoBoleto)
}

type verificaConta interface{
	Sacar(valor float64) string
}

func main() {
	// Com Ponteiros
	fmt.Println("=========== Conta da Cris ===========")
	// Cliente Cris
	clienteCris := new(clientes.Titular)
	clienteCris.CPF = "063.580.380-10"
	clienteCris.Nome = "Cris Souza"
	clienteCris.Profissao = "Médica"
	// Conta da Cris
	contaDaCris := new(contas.ContaCorrente)
	contaDaCris.Titular = clienteCris
	contaDaCris.NumeroAgencia = 2654
	contaDaCris.NumeroConta = 12456
	contaDaCris.Depositar(12503.29)
	fmt.Println(*contaDaCris.Titular)
	fmt.Println(*contaDaCris)
	fmt.Println(contaDaCris.ConsultarSaldo())

	fmt.Println("=========== Conta do Denisson ===========")
	// Cliente Denisson
	clienteDenisson := new(clientes.Titular)
	clienteDenisson.CPF = "538.240.490-90"
	clienteDenisson.Nome = "Denisson Freitas"
	clienteDenisson.Profissao = "Engenheiro DevOps"
	// Conta do Denisson
	contaDoDenisson := new(contas.ContaCorrente)
	contaDoDenisson.Titular = clienteDenisson
	contaDoDenisson.NumeroAgencia = 1222
	contaDoDenisson.NumeroConta = 35694
	contaDoDenisson.Depositar(43520.25)
	fmt.Println(*contaDoDenisson.Titular)
	fmt.Println(*contaDoDenisson)
	fmt.Println(contaDoDenisson.ConsultarSaldo())


	fmt.Println("=========== Conta do Denisson : Pagar Boleto ===========")
	PagarBoleto(contaDoDenisson, 520.25);
	fmt.Println(contaDoDenisson.ConsultarSaldo())	
}