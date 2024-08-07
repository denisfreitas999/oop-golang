# Liguagem Go

A linguagem Go, também conhecida como GoLang, é uma linguagem open source que foi criada pelo Google em 2007, e desde então é utilizada para a construção de produtos e serviços de grande escala. Atualmente a linguagem é utilizada por diversas empresas, como Uber, Twitch, Medium e Mercado livre.

Go é uma linguagem simples e produtiva de se utilizar, com foco no desenvolvimento de aplicações que necessitam de alta performance. Embora tenha sido criada para lidar com sistemas de redes e infraestrutura, Go também é bastante utilizada no mercado para:

Desenvolvimento de aplicações server-side e hospedadas em ambientes cloud;
Construção de scripts e ferramentas de automações utilizadas por times DevOps;
Construção de ferramentas de linha de comando;
Soluções de inteligência artifical e data science.

Fonte: [Formação Go](https://cursos.alura.com.br/formacao-go)

## Curso
#### Curso 02 - [Go: Orientação a Objetos](https://www.alura.com.br/curso-online-go-lang-oo)

## Referências Anteriores
[Repositório Iniciando com o Go](https://github.com/denisfreitas999/starting-golang)

## Índices
1. [Golang Orientado a Objetos](#01-golang-orientado-a-objetos)
2. [Struct](#02-struct)
3. [New e Ponteiros](#03-new-e-ponteiros)
4. [Funções e Ponteiros](#04-funções-e-ponteiros)
5. [Múltiplos Retornos](#05-múltiplos-retornos)
6. [Go Mod](#06-go-mod)
7. [Módulos e Pacotes](#07-módulos-e-pacotes-package)

## 01. Golang Orientado a Objetos

Embora Go não suporte herança tradicional como outras linguagens OOP (por exemplo, Java ou C++), ele permite a implementação de muitos conceitos de OOP através de suas próprias construções, como estruturas e interfaces.

- **Estruturas (`struct`):** Go utiliza estruturas para definir tipos de dados compostos. Estruturas são semelhantes às classes em outras linguagens OOP, permitindo agrupar campos relacionados e definir métodos associados. No entanto, não há suporte nativo para métodos incorporados diretamente nas estruturas; em vez disso, os métodos são definidos com um receptor de estrutura.

- **Métodos:** Funções associadas a tipos específicos (estruturas) usando um receptor. Os métodos permitem que você defina comportamentos que podem operar sobre as instâncias das suas estruturas.

- **Encapsulamento:** Em Go, o encapsulamento é alcançado através da convenção de nomenclatura: identificadores que começam com letras maiúsculas são exportados e acessíveis fora do pacote, enquanto aqueles que começam com letras minúsculas são privados ao pacote. Assim, você pode controlar a visibilidade de campos e métodos.

- **Composição:** Em vez de herança, Go utiliza composição para reutilizar código. Isso é feito incluindo um tipo em outro tipo (estrutura), permitindo que você aproveite o comportamento de um tipo dentro de outro. Isso oferece flexibilidade e permite a criação de tipos mais complexos sem a necessidade de uma hierarquia de herança.

- **Interfaces:** Interfaces são um conceito fundamental em Go para definir comportamentos que tipos devem implementar. Elas permitem a abstração e o polimorfismo, permitindo que diferentes tipos implementem o mesmo conjunto de métodos. Isso permite a criação de código que pode trabalhar com qualquer tipo que implemente uma interface específica, promovendo a reutilização e a flexibilidade.

**Principais Conceitos de OOP em Go:**

- **Estruturas e Métodos:** Usados para definir tipos de dados e comportamentos.
- **Encapsulamento:** Através de nomenclatura e visibilidade de campos e métodos.
- **Composição:** Reutilização de código sem herança tradicional.
- **Interfaces:** Definem comportamentos e permitem polimorfismo.

Apesar de não ter suporte para herança tradicional, Go fornece ferramentas poderosas para organizar e estruturar código de maneira eficiente e modular, utilizando conceitos de OOP adaptados à sua filosofia de design.


## 02. **Struct**

Em Go, uma **struct** é uma coleção de campos que podem conter diferentes tipos de dados. Structs são usadas para agrupar dados relacionados, facilitando a organização e manipulação dessas informações. Ao definir uma struct, você cria um novo tipo que pode ser utilizado para instanciar variáveis contendo múltiplos valores agrupados.

**Exemplo de Definição e Uso de uma Struct:**
```go
type ContaCorrente struct {
    titular       string
    numeroAgencia int
    numeroConta   int
    saldo         float64
}

func main() {
    fmt.Println(ContaCorrente{})
}
```

No exemplo acima, a struct `ContaCorrente` é definida com quatro campos: `titular`, `numeroAgencia`, `numeroConta` e `saldo`. Esses campos representam as características de uma conta corrente.

Quando uma struct é instanciada sem valores iniciais, como em `ContaCorrente{}`, seus campos assumem os **Zero Values** de seus respectivos tipos.

**Zero Value:**

O conceito de "Zero Value" em Go refere-se ao valor padrão que é atribuído automaticamente a cada campo de uma struct quando ela é inicializada sem valores. O Zero Value varia de acordo com o tipo de dado:

- Para strings: `""` (string vazia)
- Para inteiros: `0`
- Para floats: `0.0`
- Para booleanos: `false`
- Para structs: `{}`
- Para pointers, slices, maps, channels, funções e interfaces: `nil`

**Exemplo de Saída com Zero Value:**
```sh
PS E:\alura\github\curso-go-poo\src> go run .\main.go
{  0 0 0}
```

Neste exemplo, a struct `ContaCorrente` foi inicializada sem valores, resultando na saída `{  0 0 0}`. Isso ocorre porque `titular` é uma string que assume o valor vazio `""`, enquanto `numeroAgencia`, `numeroConta` e `saldo` são numéricos e assumem o valor `0`.

Esse comportamento é útil porque permite a criação de structs sem precisar inicializar todos os campos explicitamente, utilizando os Zero Values como padrões.


**Alimentando uma Struct**

Em Go, você pode inicializar e alimentar uma struct com valores específicos de duas maneiras: usando a sintaxe de inicialização nomeada ou a sintaxe posicional. Ambas permitem definir valores para os campos da struct, mas têm algumas diferenças na forma como os dados são fornecidos.

**Exemplo de Inicialização Nomeada:**
```go
contaDoDenisson := ContaCorrente{
    titular:       "Denisson Freitas",
    numeroAgencia: 1234,
    numeroConta:   54321,
    saldo:         125.50,
}
```

Neste exemplo, os valores são atribuídos a campos específicos da struct `ContaCorrente` usando a sintaxe nomeada. Essa abordagem é útil quando você deseja garantir que os valores sejam atribuídos corretamente aos campos apropriados e melhora a legibilidade do código.

**Exemplo de Inicialização Posicional:**
```go
contaDoBinho := ContaCorrente{
    "Binho Filho", // titular
    9876,         // numeroAgencia
    12345,        // numeroConta
    1000,         // saldo
}
```

No exemplo acima, os valores são atribuídos aos campos da struct `ContaCorrente` em uma ordem posicional, correspondente à ordem de declaração dos campos na struct. Essa abordagem é mais compacta, mas menos explícita, e pode ser mais suscetível a erros se a ordem dos campos não for bem compreendida.

Ambos os métodos são válidos e podem ser usados conforme a necessidade. A escolha entre eles geralmente depende da clareza desejada e da complexidade da struct.

## 03. **New e Ponteiros**

Em Go, além da inicialização direta de structs, é possível manipular e condicionar valores a elas utilizando ponteiros. Utilizar ponteiros pode ser útil quando você precisa modificar a struct em diferentes partes do código ou quando deseja evitar a cópia de grandes estruturas de dados.

**Ponteiros e a Função `new`:**

A função `new` em Go é usada para alocar memória para uma nova variável do tipo especificado e retorna um ponteiro para essa variável. Esse ponteiro pode ser utilizado para acessar e modificar a struct diretamente.

**Exemplo de Uso de Ponteiros com `new`:**
```go
// Utilizando ponteiros
var contaDaCris *ContaCorrente
contaDaCris = new(ContaCorrente)
contaDaCris.titular = "Cris"
contaDaCris.numeroAgencia = 2654
contaDaCris.numeroConta = 12456

fmt.Println(*contaDaCris)
```

No exemplo acima:

1. **Declaração do Ponteiro:**
    ```go
    var contaDaCris *ContaCorrente
    ```
    Aqui, `contaDaCris` é declarado como um ponteiro para uma `ContaCorrente`.

2. **Alocação de Memória com `new`:**
    ```go
    contaDaCris = new(ContaCorrente)
    ```
    A função `new(ContaCorrente)` aloca memória para uma nova `ContaCorrente` e retorna um ponteiro para essa memória. O ponteiro é então atribuído à variável `contaDaCris`.

3. **Atribuição de Valores:**
    ```go
    contaDaCris.titular = "Cris"
    contaDaCris.numeroAgencia = 2654
    contaDaCris.numeroConta = 12456
    ```
    Os campos da struct são acessados e modificados através do ponteiro. Quando você usa o ponteiro `contaDaCris`, você está manipulando diretamente a instância da struct alocada na memória.

4. **Impressão da Struct:**
    ```go
    fmt.Println(*contaDaCris)
    ```
    O operador `*` é usado para desreferenciar o ponteiro e acessar o valor da struct. Sem esse operador, você estaria imprimindo o endereço de memória.

Usar ponteiros pode ser benéfico para evitar a cópia de structs grandes, reduzir o uso de memória e permitir que você modifique a struct original sem criar cópias adicionais. No entanto, é importante gerenciar ponteiros com cuidado para evitar problemas como referências nulas ou vazamentos de memória.

## 04. **Funções e Ponteiros**

Em Go, você pode definir métodos para structs que utilizam ponteiros como receivers. Isso é especialmente útil quando você precisa modificar os campos da struct dentro de uma função, pois o uso de ponteiros permite que a função altere diretamente os dados originais em vez de trabalhar com uma cópia.

**Exemplo de Uso de Funções com Ponteiros:**
```go
func main() {
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
```

Neste exemplo:

1. **Declaração da Struct e Inicialização:**
    A struct `ContaCorrente` é instanciada usando a função `new`, e os campos da struct são preenchidos com valores.

2. **Uso de Ponteiros como Receivers:**
    A função `sacar` é definida com um receiver que é um ponteiro para `ContaCorrente`:
    ```go
    func (c *ContaCorrente) sacar(valorDoSaque float64) string
    ```
    Isso significa que quando `sacar` é chamado, ele pode modificar diretamente os campos da struct original, como o campo `saldo`.

3. **Modificação dos Campos:**
    Dentro da função `sacar`, o valor do saque é subtraído diretamente do saldo da conta, utilizando o ponteiro `c`. Como `c` é um ponteiro, a modificação afeta a struct original em memória, e não uma cópia.

4. **Chamada da Função:**
    No `main`, o método `sacar` é chamado usando `contaDaCris.sacar(500)`. Após a execução da função, o saldo da conta é reduzido, refletindo a operação de saque.

Usar ponteiros como receivers em funções é uma prática comum em Go quando você precisa alterar o estado de uma struct. Isso permite que as alterações feitas dentro da função sejam persistentes e reflitam diretamente na instância original da struct.

## 05. **Múltiplos Retornos**

Em Go, as funções e métodos podem retornar múltiplos valores. Isso é útil em situações onde você deseja retornar tanto um resultado primário quanto informações adicionais, como um status ou um erro. O exemplo a seguir demonstra como utilizar múltiplos retornos em um método que realiza depósitos em uma conta corrente.

**Exemplo de Método com Múltiplos Retornos:**
```go
func (c *ContaCorrente) Depositar(valorDoDeposito float64) (string, float64) {
    podeDepositar := valorDoDeposito > 0

    if podeDepositar {
        c.saldo += valorDoDeposito
        return "Depósito realizado com sucesso!", c.saldo
    } else {
        return "Não foi possível realizar o depósito. Entre com um valor válido.", c.saldo
    }
}
```

Neste exemplo:

1. **Definição do Método `Depositar`:**
    O método `Depositar` recebe um valor de depósito como argumento e retorna duas informações:
    - Uma mensagem de status (`string`)
    - O saldo atual da conta (`float64`)

2. **Verificação e Atualização:**
    O método verifica se o valor do depósito é positivo. Se for, o saldo da conta é atualizado e uma mensagem de sucesso é retornada junto com o novo saldo. Caso contrário, uma mensagem de erro é retornada junto com o saldo atual, que não é alterado.

**Exemplo de Uso de Múltiplos Retornos:**
```go
fmt.Println(contaDaCris.Depositar(1000))  // Saída: "Depósito realizado com sucesso!", 2000
fmt.Println(*contaDaCris)                // Saída: {Cris 2654 12456 2000}

status, saldo := contaDaCris.Depositar(200)
fmt.Println(status)  // Saída: "Depósito realizado com sucesso!"
fmt.Println(saldo)   // Saída: 2200
```

Neste trecho de código:

- **Chamada com `fmt.Println`:** 
    A primeira chamada a `Depositar` é usada diretamente dentro de `fmt.Println`, que imprime os dois valores retornados.

- **Atribuição a Variáveis:**
    Na segunda chamada, os valores retornados por `Depositar` são capturados em variáveis separadas, `status` e `saldo`. Isso permite um maior controle sobre como cada valor é manipulado e exibido posteriormente.

O uso de múltiplos retornos em Go permite que funções e métodos sejam mais expressivos e forneçam informações adicionais sem a necessidade de estruturas mais complexas como tuplas ou objetos. Isso é especialmente útil para retornar um resultado primário e um status ou erro associado.

## 06. **Go Mod**

O `go mod` é uma ferramenta usada no Go para gerenciar dependências de projetos, introduzida a partir da versão 1.11 do Go. Com o conceito de módulos, o `go mod` facilita o gerenciamento de bibliotecas e pacotes externos, eliminando a necessidade de configurar o GOPATH e proporcionando uma abordagem mais moderna e eficiente para o controle de dependências.

**O que é `go.mod`?**

O arquivo `go.mod` é o arquivo de configuração principal para um módulo Go. Ele contém informações sobre o módulo, incluindo:
- O nome do módulo, que geralmente corresponde ao caminho do repositório onde o código está armazenado.
- As versões das dependências utilizadas pelo projeto.
- Outras informações de configuração do módulo, como os requisitos de versão do Go.

**Como criar o `go.mod`?**

Para criar um arquivo `go.mod`, você precisa inicializar um módulo no seu projeto. Isso pode ser feito utilizando o comando `go mod init`. Siga os passos abaixo:

1. **Abra um terminal e navegue até o diretório do seu projeto.**
2. **Execute o comando `go mod init` seguido do caminho do módulo (geralmente o caminho do repositório).**
    ```sh
    go mod init exemplo.com/meuprojeto
    ```
3. **O Go criará um arquivo `go.mod` no diretório do projeto com a configuração inicial do módulo.**

**Benefícios do `go.mod`:**

- **Gerenciamento Simplificado de Dependências:** O `go mod` facilita o gerenciamento das versões das dependências, permitindo que você especifique versões exatas e bloqueie versões específicas para garantir a consistência do projeto.
- **Isolamento do GOPATH:** Com o uso de módulos, não é mais necessário configurar o GOPATH, tornando o gerenciamento de projetos mais simples e menos propenso a erros.
- **Reprodutibilidade e Consistência:** O arquivo `go.mod` e o arquivo `go.sum` (que registra as somas de verificação das dependências) ajudam a garantir que todos os desenvolvedores e ambientes de build utilizem as mesmas versões das dependências, promovendo a consistência do projeto.
- **Facilidade na Importação de Pacotes:** Com o `go.mod`, você pode importar pacotes diretamente sem precisar ajustar o GOPATH, facilitando a inclusão de bibliotecas externas no seu projeto.

**Exemplo de `go.mod`:**
```go
module exemplo.com/meuprojeto

go 1.18

require (
    github.com/some/dependency v1.2.3
    github.com/another/dependency v4.5.6
)
```

No exemplo acima, o arquivo `go.mod` define o módulo como `exemplo.com/meuprojeto`, especifica a versão do Go e lista as dependências do projeto com suas versões correspondentes.

Utilizar `go mod` é uma prática recomendada para projetos Go modernos, pois oferece uma abordagem mais robusta e flexível para gerenciamento de dependências.

## 07. **Módulos e Pacotes (package)**

Em Go, um **package** (ou pacote) é uma maneira de organizar o código em módulos reutilizáveis. Um pacote é um conjunto de arquivos Go que estão localizados no mesmo diretório e têm o mesmo nome de pacote. Pacotes permitem que você agrupe funções, tipos e variáveis relacionadas, facilitando a modularização e a reutilização do código.

**Como funciona?**

Cada arquivo Go em um diretório faz parte de um pacote. O nome do pacote é declarado no início de cada arquivo com a palavra-chave `package`. Quando você importa um pacote em outro arquivo, pode acessar as funções, tipos e variáveis exportados pelo pacote.

**Regras de Visibilidade:**

- **Público:** Identificadores (como funções, métodos, variáveis e tipos) que começam com uma letra maiúscula são exportados e, portanto, públicos. Isso significa que eles podem ser acessados a partir de outros pacotes.
- **Privado:** Identificadores que começam com uma letra minúscula não são exportados e, portanto, privados. Esses identificadores só podem ser acessados dentro do próprio pacote.

**Como utilizar?**

Para utilizar um pacote, você precisa importá-lo usando a palavra-chave `import`. O caminho do pacote especifica onde ele está localizado. Você pode então acessar os identificadores exportados do pacote.

**Exemplo:**

**Arquivo `contas.go` no pacote `contas`:**
```go
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
        return "Depósito realizado com sucesso!", c.Saldo
    } else {
        return "Não foi possível realizar o depósito. Entre com um valor válido.", c.Saldo
    }
}

func (c *ContaCorrente) Transferir(valorDaTransferencia float64, contaDestino *ContaCorrente) string {
    podeTransferir := valorDaTransferencia > 0 && c.Saldo > valorDaTransferencia

    if podeTransferir {
        c.Saldo -= valorDaTransferencia
        contaDestino.Depositar(valorDaTransferencia)
        return "Transferência realizada com sucesso!"
    } else {
        return "Ocorreu um erro na transferência. Por favor, verifique os valores inseridos."
    }
}
```

**Arquivo `main.go` no pacote `main`:**
```go
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

    // Realizando saques na conta da Cris
    fmt.Println(contaDaCris.Sacar(500))
    fmt.Println(*contaDaCris)
    fmt.Println(contaDaCris.Sacar(1000))
    fmt.Println(*contaDaCris)
    fmt.Println(contaDaCris.Depositar(1000))
    fmt.Println(*contaDaCris)
    status, saldo := contaDaCris.Depositar(200)
    fmt.Println(status)
    fmt.Println(saldo)

    // Criando a conta do Denisson
    contaDoDenisson := new(contas.ContaCorrente)
    contaDoDenisson.Titular = "Denisson Freitas"
    contaDoDenisson.NumeroAgencia = 1222
    contaDoDenisson.NumeroConta = 35694
    contaDoDenisson.Saldo = 1200
    fmt.Println(*contaDoDenisson)

    // Testando transferências
    fmt.Println("========== Valores iniciais nas contas dos usuários =========")
    fmt.Println(*contaDoDenisson)
    fmt.Println(*contaDaCris)
    fmt.Println("========== Denisson transfere 200 para Cris =================")
    contaDoDenisson.Transferir(200, contaDaCris)
    fmt.Println("========== Valores nas contas após a transferência ==========")
    fmt.Println(*contaDoDenisson)
    fmt.Println(*contaDaCris)
    fmt.Println("========== Cris transfere 400 para Denisson =================")
    contaDaCris.Transferir(400, contaDoDenisson)
    fmt.Println("========== Valores nas contas após a transferência ==========")
    fmt.Println(*contaDoDenisson)
    fmt.Println(*contaDaCris)
}
```

**Explicação:**

- **Pacote `contas`:** Define o tipo `ContaCorrente` com métodos para sacar, depositar e transferir valores. Os métodos são exportados porque começam com letras maiúsculas, permitindo que sejam usados em outros pacotes.
- **Pacote `main`:** Importa o pacote `contas` e utiliza suas funções e tipos para criar e manipular instâncias de `ContaCorrente`. O código realiza operações de saque, depósito e transferência.

Pacotes são fundamentais para a organização e modularização de projetos Go, ajudando a separar responsabilidades e a manter o código mais limpo e gerenciável.
