package main

import (
	"curso-go-poo/pkg/contas"
	"fmt"
)

func main() {
	contaDaCris := new(contas.ContaCorrente)
	contaDaCris.Titular = "Cris Souza"
	contaDaCris.NumeroAgencia = 2654
	contaDaCris.NumeroConta = 12456
	contaDaCris.Saldo = 1000
	fmt.Println(*contaDaCris)

	// Realizando Saques na Conta da Cris
	fmt.Println(contaDaCris.Sacar(500))
	fmt.Println(*contaDaCris)
	fmt.Println(contaDaCris.Sacar(1000))
	fmt.Println(*contaDaCris)
	fmt.Println(contaDaCris.Depositar(1000))
	fmt.Println(*contaDaCris)
	status, Saldo := contaDaCris.Depositar(200)
	fmt.Println(status)
	fmt.Println(Saldo)

	// Criando a conta do Denisson
	contaDoDenisson := new(contas.ContaCorrente)
	contaDoDenisson.Titular = "Denisson Freitas"
	contaDoDenisson.NumeroAgencia = 1222
	contaDoDenisson.NumeroConta = 35694
	contaDoDenisson.Saldo = 1200
	fmt.Println(*contaDoDenisson);

	// Testando Transferências
	fmt.Println("========== Valores iniciais nas contas dos usuários =========")
	fmt.Println(*contaDoDenisson);
	fmt.Println(*contaDaCris)
	fmt.Println("========== Denisson transfere 200 para Cris =================")
	contaDoDenisson.Transferir(200, contaDaCris)
	fmt.Println("========== Valores nas contas após a transferência ==========")
	fmt.Println(*contaDoDenisson);
	fmt.Println(*contaDaCris)
	fmt.Println("========== Cris transfere 400 para Denisson =================")
	contaDaCris.Transferir(400, contaDoDenisson)
	fmt.Println("========== Valores nas contas após a transferência ==========")
	fmt.Println(*contaDoDenisson);
	fmt.Println(*contaDaCris)
}