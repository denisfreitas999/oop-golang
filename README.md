# Liguagem Go

A linguagem Go, também conhecida como GoLang, é uma linguagem open source que foi criada pelo Google em 2007, e desde então é utilizada para a construção de produtos e serviços de grande escala. Atualmente a linguagem é utilizada por diversas empresas, como Uber, Twitch, Medium e Mercado livre.

Go é uma linguagem simples e produtiva de se utilizar, com foco no desenvolvimento de aplicações que necessitam de alta performance. Embora tenha sido criada para lidar com sistemas de redes e infraestrutura, Go também é bastante utilizada no mercado para:

Desenvolvimento de aplicações server-side e hospedadas em ambientes cloud;
Construção de scripts e ferramentas de automações utilizadas por times DevOps;
Construção de ferramentas de linha de comando;
Soluções de inteligência artifical e data science.

Fonte: [Formação Go](https://cursos.alura.com.br/formacao-go)

## Cursos
#### Curso 02 - [Go: Orientação a Objetos](https://www.alura.com.br/curso-online-go-lang-oo)

## Referências Anteriores
[Repositório Iniciando com o Go](https://github.com/denisfreitas999/starting-golang)

## Indices
1. [Introdução](#01-golang-orientado-a-objetos)

## 01. Golang Orientado a Objetos

01. **Struct**

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

2. **New e Ponteiros**

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

3. **Funções e Ponteiros**

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

