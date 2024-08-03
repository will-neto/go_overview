package main

// Para um arquivo em Go fazer parte do programa deve ser declarado no topo do arquivo o comando "package [nome-pacote]"
// O nome do pacote pode ser qualquer um. Porém, o Go possui um pacote predefinido chamado "main" (usado acima). Que indica que o arquivo é executável.
// Através do comando "go run [nome-arquivo]" é possível compilar + executar o programa .go que seja atribuido com o "package main" e que possua a função main() declarada
// No módulo deve ter apenas um arquivo com a função main, não é permitido dois arquivos no mesmo módulo com o pacote main e função main() atribuida

import (
	"fmt" // Biblioteca nativa do Go para trabalhar com formatação e impressão de dados
)

func main() {
	fmt.Println("Exemplo de execução de programa do package main!")
}
