package main

import (
	"fmt"
	// A importação de um pacote deve seguir a estrutura seguinte: [nome-modulo]/[path-diretorio-pacote]
	// No exemplo abaixo é passado o "pacotes_compartilhamento", como [nome-modulo], pois é o nome declarado no arquivo go.mod
	// Em seguida passamos "/auxiliar", pois é o diretório onde está nosso "package auxiliar"
	"pacotes_compartilhamento/auxiliar"
	// Podemos também trabalhar com pacotes que estejam mais de um nível de path de subdiretório
	"pacotes_compartilhamento/exemplo_subdiretorio/subdiretorio"
	// Em cenários onde o nome do pacote não é o mesmo nome do diretório que ele está (devido a flexibilidade que possuimos de organiszação)
	//		, devemos seguir a estrutura seguinte: [nome-pacote] [nome-modulo]/[path-diretorio-pacote]
	// No exemplo abaixo é passado "exemplo" como [nome-pacote], seguido por "pacotes_compartilhamento" como [nome-modulo] e "/exemplo_pacote_nome_diferente_diretorio" como path da localização do pacote
	exemplo "pacotes_compartilhamento/exemplo_pacote_nome_diferente_diretorio"
	// O exemplo acima também trabalha com mais de um nível de path de subdiretório
)

func main() {
	fmt.Println("Pacote Main")
	auxiliar.ExibirMensagemMetodoPublico()
	subdiretorio.ExibirMensagem()
	exemplo.ExibirMensagemPacoteExemplo()
}
