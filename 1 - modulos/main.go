package main

// Embora os pacotes em Go possam existir sem módulos, é recomendável que sejam gerenciados através de módulos
// Módulos são conjuntos de pacotes (packages) que são gerenciados de maneira agrupada
// O nome do módulo é definido através do arquivo "go.mod", onde também podem ser especificadas as dependências
// O arquivo de módulo (go.mod) pode ser criado manualmente ou através do comando "go mod init [nome-modulo]" no diretório raiz do módulo, também sendo recomendado que o nome do modulo seja do diretório raiz
// O arquivo de módulo (go.mod) deve estar no diretório raiz do módulo
// Após a criação do arquivo de módulo (go.mod), temos a flexibilidade de criar os pacotes no diretório raiz ou criando subdiretórios (dentro do diretório do módulo) para organizar o código da nossa maneira

import "fmt"

func main() {
	fmt.Println("Escrevendo do arquivo main()")
}

// Após a criação do módulo, podemos executar o comando "go build" no diretório raiz do módulo para compilação dos arquivos do módulo
// A execução do comando "go build" por padrão gera um executável (.exe), o que requer um pacote "package main" (que é um pacote especial do Go e que pode ser atribuido a um arquivo, no caso, este com a definição do método main())
// Porém, o comando aceita a geração de outros tipos de arquivos através da flag -buildmode => go build -buildmode=[tipo-arquivo]
// Alguns tipos de arquivos, como a geração de biblioteca (para utilização por outros programas), talvez não necessitem de um arquivo com package main e função main() para sua execução
