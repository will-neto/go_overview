package auxiliar

// Podemos definir qualquer nome de pacote para os arquivos que formos criando, não é obrigatório que o nome do pacote seja o nome do diretório onde ele pertence
// Não é possível existir dois pacotes diferentes no mesmo diretório (exemplo, temos o package auxiliar declarado acima o que implica na obrigatoriedade de mantermos o mesmo package auxiliar em todos os arquivos do diretório)
// Pacotes ajudam a organizar e modularizar o código, permitindo que você agrupe funções, tipos e variáveis relacionados
// Pacotes podem ser fracionados em mais de um arquivo (auxiliar2.go) e enxergar as declarações do outro arquivo, pois no final são unitários (pertencem ao mesmo pacote). Esse fracionamento entre arquivos pode ajudar na organização separada do código

import (
	"fmt"
)

// Métodos publicos (que podem ser acessados externamente por outros pacotes/modulos) devem seguir a nomenclatura Pascal Case
// Esse método pode ser usado no arquivo "main.go", pois importamos o pacote (package auxiliar) por lá e temos visibilidade por ele ser público
func ExibirMensagemMetodoPublico() {
	fmt.Println("Exibindo mensagem através do arquivo auxiliar.go")

	// Como comentado acima, podemos chamar métodos de outros arquivos que estejam no mesmo pacote (no caso, o arquivo auxiliar2.go pertence ao package auxiliar)
	exibirMensagemMetodoPrivado()
}

// Método privado que não pode ser chamado por fora do "package auxiliar"
func metodoPrivado() {
	fmt.Println("Exibindo mensagem privada do método privado declarado no arquivo auxiliar.go")
}
