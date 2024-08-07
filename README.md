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

    