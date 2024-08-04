package auxiliar

import (
	"fmt"
)

// Métodos privados (que podem não podem ser acessados externamente, apenas dentro do mesmo pacote) devem seguir a nomenclatura Camel Case
// Esse método não pode ser usado no arquivo "main.go" devido o pacote "package main" não ter visibilidade
func exibirMensagemMetodoPrivado() {
	fmt.Println("Exibindo mensagem privada do método privado declarado no arquivo auxiliar2.go")

	// Chamando método privado que pode ser acessa por este arquivo (auxiliar2.go) devido pertencem ao mesmo pacote do auxiliar.go (no caso, package auxiliar)
	metodoPrivado()
}
