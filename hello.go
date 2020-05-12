package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const vezesMonitoramento = 5
const delayMonitoramento = 5 * time.Second

func main() {
	exibeIntroducao()
	for {
		exibeMenu()

		//var comando string = "conteudo" | nessa declaração de variavel é necessário que o programador informe o tipo da variavel, podendo ser 'string, int, float32, float64, etc', caso queira pode ser passado ser a identificação do tipo, o go irá reconhecer nesse caso o tipo no momento que ele receba o conteudo
		//var comando string | se eu declarar sem o conteudo, ele atribuirá o valor vazio do tipo, ou seja, "" para string, 0 para int e assim por diante
		comando := leComando() // posso declarar uma variavel utilizando ':=', nesse formato a variavel já indentifica qual o tipo dela automaticamente

		switch comando {
		case 1:
			iniciarMonitoramento()
		case 2:
			imprimeLogs()
		case 0:
			fmt.Println("Saindo do programa...")
			os.Exit(0)
		default:
			fmt.Println("Comando não reconhecido!")
			os.Exit(-1)
		}
	}

}

// estrutura de função básica
func exibeIntroducao() {
	nome := "Marcos"
	versao := 1.0
	fmt.Println("Olá, sr.", nome)
	fmt.Println("Este programa está na versão", versao)
}

func exibeMenu() {
	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("0 - Sair do Programa")
}

// estrutura de função com retorno especificado, no caso retorno 'int'
func leComando() int {
	var comandoLido int
	fmt.Scan(&comandoLido)
	fmt.Println("O comando escolhido foi:", comandoLido)
	return comandoLido
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")
	// sites := []string{"https://alura.com.br", "https://baixefacil.com.br", "https://apkeasy.me"}
	sites := leSitesDoArquivo()

	for i := 0; i < vezesMonitoramento; i++ {
		for _, site := range sites {
			testaSite(site)
		}
		time.Sleep(delayMonitoramento)
		fmt.Println("")
	}

	fmt.Println("")

}

// estrutura de função com parametros espeficiado, é necessário indicar o tipo
func testaSite(site string) {
	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
		registraLog(site, true)
	} else {
		fmt.Println("Site:", site, "está com problemas. Status Code:", resp.StatusCode)
		registraLog(site, false)
	}
}

func leSitesDoArquivo() []string {
	var sites []string
	arquivo, err := os.Open("sites.txt")
	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	leitor := bufio.NewReader(arquivo)

	for {
		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)

		sites = append(sites, linha)

		if err == io.EOF {
			break
		}
	}

	arquivo.Close()
	return sites
}

func registraLog(site string, status bool) {
	arquivo, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666) // 'OpenFile', 'Open' da biblioteca 'os' abre os arquivos para a leitura, entretanto é necessário fechar
	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + " - online: " + strconv.FormatBool(status) + "\n")

	arquivo.Close()
}

func imprimeLogs() {
	fmt.Println("Exibindo Logs...")
	arquivo, err := ioutil.ReadFile("log.txt") //a função 'ReadFile' da biblioteca 'ioutil' já realiza o fechamento, ela manda a leitura em byte e fecha a leitura do arquivo, não necessitando fechar manualmente como na biblioteca 'os'
	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}
	fmt.Println(string(arquivo))
}
