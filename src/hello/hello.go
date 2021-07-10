package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
	"io/ioutil"
	"bufio"
	"strconv"
)

const nMonitor = 5
const timeDelay = 5

func main() {

	intro()

	for {
		showMenu()

		command := readCommand()

		// if command == 1 {
		// 	fmt.Println("Monitoring...")
		// } else if command == 2 {
		// 	fmt.Println("Showing logs...")
		// } else if command == 0 {
		// 	fmt.Println("Exiting program...")
		// } else {
		// 	fmt.Println("Invalid command!")
		// }

		switch command {
		case 1:
			startMonitoring()
		case 2:
			fmt.Println("Showing logs...")
			printLogs()
		case 0:
			fmt.Println("Exiting program...")
			os.Exit(0)
		default:
			fmt.Println("Invalid command!")
			os.Exit(-1)

		}
	}
}

func intro() {
	//var name string = "Amorim"
	//var version float64 = 1.1
	version := 1.1
	name := "Amorim"
	fmt.Println("Hello, Mr.", name)
	fmt.Println("This program is on version", version)
	fmt.Println()
}

func showMenu() {
	fmt.Println("1- Start monitoring")
	fmt.Println("2- Show logs")
	fmt.Println("0- Exit program")
}

func readCommand() int {
	command := 0
	fmt.Scan(&command)
	fmt.Println("The chosen command was ", command)
	return command
}

func startMonitoring() {
	fmt.Println("Monitoring...")

	sites := readSites()

	//sites := []string{
	//	"https://bemol.com.br",
	//	"https://alura.com.br",
	//	"https://youtube.com",
	//	"https://google.com.br"}

	for i := 0;  i < nMonitor; i++{
		for i, j := range sites {
			fmt.Println("Test",i,"site:", j)
			testSite(j)
		}
		time.Sleep(timeDelay * time.Second)
	}
}

func testSite(site string) {
	response, err := http.Get(site)
	if err != nil {
		fmt.Println("Error detected!", err)
	}

	if response.StatusCode == 200 {
		fmt.Println(site, "está no ar.")
		saveLog(site, true)
	} else {
		fmt.Println(site, "está fora do ar. Status code:", response.StatusCode)
		saveLog(site, false)
	}
}

func readSites()[]string {

	var sites = []string

	//file, err := os.Open('sites.txt')
	file, err := ioutil.readFile('site.txt')
	if err != nil {
		fmt.Println("Error detected!", err)

	}

	reader := bufio.NewReader(file)


	for {
		line, err := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		sites = append(sites, line)

		if err == io.EOF {
			break
		}

	}
	return sites
}

func saveLog(site string, status bool){
	file, err := os.OpenFile('log.txt', os.O_RDWR|os.O_CREATE|os.O_APPEND,0666)
	if err != nil {
		fmt.Println(err)
	}
	file.WriteString(time.Now().Format("02/01/2006 15:04:05")+ " | " + site + " | online: "+ strconv.FormatBool(status)+"\n")
	file.Close()
}

func printLogs() {
	file, err := ioutil.ReadFile("log.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(file))
}

// Go code from Alura course
/*
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

const monitoramentos = 2
const delay = 5

func main() {
    exibeIntroducao()
    for {
        exibeMenu()
        comando := leComando()

        switch comando {
        case 1:
            iniciarMonitoramento()
        case 2:
            //Chamando aqui
            fmt.Println("Exibindo Logs...")
            imprimeLogs()
        case 0:
            fmt.Println("Saindo do programa")
            os.Exit(0)
        default:
            fmt.Println("Não conheço este comando")
            os.Exit(-1)
        }
    }
}

func exibeIntroducao() {
    nome := "Douglas"
    versao := 1.2
    fmt.Println("Olá, sr.", nome)
    fmt.Println("Este programa está na versão", versao)
}

func exibeMenu() {
    fmt.Println("1- Iniciar Monitoramento")
    fmt.Println("2- Exibir Logs")
    fmt.Println("0- Sair do Programa")
}

func leComando() int {
    var comandoLido int
    fmt.Scan(&comandoLido)
    fmt.Println("O comando escolhido foi", comandoLido)
    fmt.Println("")

    return comandoLido
}

func iniciarMonitoramento() {
    fmt.Println("Monitorando...")

    sites := leSitesDoArquivo()

    for i := 0; i < monitoramentos; i++ {
        for i, site := range sites {
            fmt.Println("Testando site", i, ":", site)
            testaSite(site)
        }
        time.Sleep(delay * time.Second)
        fmt.Println("")
    }

    fmt.Println("")
}

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

func imprimeLogs() {

    arquivo, err := ioutil.ReadFile("log.txt")

    if err != nil {
        fmt.Println("Ocorreu um erro:", err)
    }

    fmt.Println(string(arquivo))
}

func registraLog(site string, status bool) {

    arquivo, err := os.OpenFile("log.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)

    if err != nil {
        fmt.Println("Ocorreu um erro:", err)
    }
    arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + " - online: " + strconv.FormatBool(status) + "\n")

    arquivo.Close()
}
*/
