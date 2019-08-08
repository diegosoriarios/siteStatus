package main

import ("bufio"
		"fmt"
		"io"
		"io/ioutil"
		"net/http"
		"os"
		"strconv"
		"strings"
		"time"
)

const DELAY = 3

func main(){
	for {
		menu()

		switch comando() {
			case 0:
				os.Exit(0)
			case 1:
				iniciarMonitoramento()
			case 2:
				imprimeLogs()
			default:
				fmt.Println("Escolha uma opção valida!")
		}
	}
}

func menu() {
	nome := "Diego"

	fmt.Println("Hello,", nome)
	
	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("0 - Sair do Programa")
}

func comando() int {
	var comando int
	fmt.Scanf("%d", &comando)
	return comando
}

func iniciarMonitoramento() {

	sites := leArquivo()

	for _, site := range sites {
		time.Sleep(DELAY * time.Second)
		resp, err := http.Get(site)

		if err != nil {
			fmt.Println("Ocorreu um erro", err)
		}

		if resp.StatusCode == 200 {
			fmt.Println("Site", site, "foi carregado com sucesso")
		} else {
			fmt.Println("Site", site, "está com problema. Status erro", resp.StatusCode)
		}
		registraLog(site, resp.StatusCode)
	}

	fmt.Println("")
}

func leArquivo() []string {
	var sites []string
	arquivo, err := os.Open("sites.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro", err)
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

func registraLog(site string, status int) {
	arquivo, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println("Erro", err)
	}

	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + " - " + strconv.Itoa(status) + "\n")

	arquivo.Close()
}

func imprimeLogs() {
	arquivo, err := ioutil.ReadFile("log.txt")

	if err != nil {
		fmt.Println("Erro", err)
	}

	fmt.Println(string(arquivo))
}