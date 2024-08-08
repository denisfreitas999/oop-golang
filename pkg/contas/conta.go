package contas

type Conta interface {
	Sacar(valor float64) string
	Depositar(valor float64) (string, float64)
	Transferir(valor float64, contaDestino Conta) string
	ConsultarSaldo() string
}
