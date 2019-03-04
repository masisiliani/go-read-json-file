package main

//Declaração dos pacotes utilziados nesse arquivo.
import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

//Struct Book. utilizada na nossa aplicação
type Book struct {
	Name      string `json:"name"`
	Publisher string `json:"publisher"`
	Author    string `json:"author"`
	Link      string `json:"link"`
}

//Função inicial da aplicação. Para um código mais clean (limpo) é bacana deixar apenas as chamadas de outras
//funções aqui, pois assim o código fica mais legível
func main() {

	//Abrindo o arquivo JSON. O parâmetro da func os.Open é formado pelo [caminho + o nome do arquivo]
	//que queremos abrir. Nesse caso como está no diretório raiz, colocamos apenas o nome do arquivo.
	jsonFile, err := os.Open(`book.json`)

	//Trecho verifica se houve erro na abertura do arquivo.
	if err != nil {
		//Caso tenha tido erro, ele é apresentado na tela
		fmt.Println(err)
	}

	//Sempre que possivel é importante dar 'dispose' ou 'close' nas variáveis, para não consumir memória.
	defer jsonFile.Close()

	//Aqui o arquivo é convertido para uma variável array de bytes, através do pacote "io/ioutil"
	byteValueJSON, err := ioutil.ReadAll(jsonFile)

	//Trecho verifica se houve erro na abertura do arquivo.
	//Vamos utilizar essa mesma instrução no decorrer do programa.
	if err != nil {
		//Caso tenha tido erro, ele é apresentado na tela
		fmt.Println(err)
	}

	//Declaração abreviada de um objeto do tipo Book
	objBook := Book{}

	//Conversão da variável byte em uma struct Book
	err = json.Unmarshal(byteValueJSON, &objBook)
	if err != nil {
		fmt.Println(err)
	}

	//Apresentação na tela do campo Name
	fmt.Println()
	fmt.Println("Nome do livro:", objBook.Name)

	//Alteração do campo Name
	objBook.Name = "New book name 2"

	//Apresentação no console do novo valor do campo Name
	fmt.Println("Novo nome do livro:", objBook.Name)
	fmt.Println()

	//Aqui vamos converter nosso objBook com o nome alterado em bytes
	byteValueJSON, err = json.Marshal(objBook)
	if err != nil {
		fmt.Println(err)
	}

	//Por fim vamos salvar em um arquivo JSON alterado.
	err = ioutil.WriteFile("newBook.json", byteValueJSON, 0644)
	if err != nil {
		fmt.Println(err)
	}

}
