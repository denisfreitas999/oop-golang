package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func main() {
	contaDoDenisson := ContaCorrente{titular: "Denisson Freitas", numeroAgencia: 1234, numeroConta: 54321, saldo: 125.50}
	contaDoBinho := ContaCorrente{"Binho Filho", 9876, 12345, 1000}

	fmt.Println(contaDoDenisson, contaDoBinho)
}
