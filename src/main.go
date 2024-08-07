package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func main() {
	// contaDoDenisson := ContaCorrente{titular: "Denisson Freitas", numeroAgencia: 1234, numeroConta: 54321, saldo: 125.50}
	// contaDoBinho := ContaCorrente{"Binho Filho", 9876, 12345, 1000}

	// fmt.Println(contaDoDenisson)
	// fmt.Println(contaDoBinho)
	// sacar(500, contaDoBinho);
	// fmt.Println(contaDoBinho)
	
	// Utilizando ponteiros
	var contaDaCris *ContaCorrente
	contaDaCris = new(ContaCorrente)
	contaDaCris.titular = "Cris"
	contaDaCris.numeroAgencia = 2654
	contaDaCris.numeroConta = 12456
	contaDaCris.saldo = 1000
	fmt.Println(*contaDaCris)

	// Realizando Saque
	fmt.Println(contaDaCris.sacar(500))
	fmt.Println(*contaDaCris)
	fmt.Println(contaDaCris.sacar(1000))
	fmt.Println(*contaDaCris)
}

func (c *ContaCorrente) sacar(valorDoSaque float64) string {

	podeSacar := valorDoSaque <= c.saldo && valorDoSaque > 0

	if podeSacar {
		c.saldo -= valorDoSaque
		return "Saque realizado com sucesso!"
	} else {	
		return "Saldo insuficiente."
	}
}