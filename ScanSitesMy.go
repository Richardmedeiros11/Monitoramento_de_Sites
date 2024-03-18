package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

var i int

func main() {
	for {
		cabecalho()
		comando := opcoes()

		switch comando {
		case 1:
			monitoramento()
		case 2:
			exibir_logs()
		case 0:
			saida()
		default:
			fmt.Println("Error")
			os.Exit(-1)
		}
	}
}

func monitoramento() {
	fmt.Println("Monitorando...")
	sites := leSiteDoArquivo()
	for i := range sites {
		testa_site(sites[i])
	}
	fmt.Println("\n\n")
}

func testa_site(site string) {
	fmt.Println("Site a ser monitorado> ", site)
	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("Ocorreu um erro na conexão: ", err)
	}
	i++
	if resp.StatusCode == 200 {
		fmt.Println(i, "-> Status Code: 200 OK")
		fmt.Println(i, "-> Site em funcionamento! :)")
		resp_status_logs := string(resp.Status)
		logs(site, resp_status_logs)
	} else {
		fmt.Println(i, "-> Status Code: ", resp.StatusCode)
		fmt.Println(i, "-> Erro! :(")
	}
	time.Sleep(5 * time.Second)
}

func logs(site string, status string) {
	arquivo, err := os.OpenFile("logs.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Error :(  >", err)
	}
	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05 - ") + site + " " + status + "\n")

	fmt.Println(arquivo)
	arquivo.Close()
}
func saida() {
	fmt.Println("Saindo...")
	time.Sleep(3 * time.Second)
	os.Exit(0)
}

func cabecalho() {
	version := 0.1
	fmt.Println("- ScanSitesMy -")
	fmt.Println("Version: ", version)
	fmt.Println("1 - Iniciar monitoramento")
	fmt.Println("2 - Exibir os logs")
	fmt.Println("0 - Sair")
}
func opcoes() int {
	var comandoLido int
	fmt.Scan(&comandoLido)
	return comandoLido
}

func leSiteDoArquivo() []string {
	var sites []string
	arquivo, err := os.Open("configuracoes.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro: ", err)
		fmt.Println(":(")
	}

	leitor := bufio.NewReader(arquivo)
	for {
		linha, err := leitor.ReadString('\n')
		if err != io.EOF && err != nil {
			fmt.Println("Ocorreu um erro: ", err)
			fmt.Println(":(")
		}
		if err == io.EOF {
			break
		}
		linha = strings.TrimSpace(linha)
		sites = append(sites, linha)
	}
	fmt.Println(sites)
	arquivo.Close()
	return sites
}

func exibir_logs() {
	arquivo, err := ioutil.ReadFile("logs.txt")
	if err != nil {
		fmt.Println("Erro :( > ", err)
	}
	fmt.Println(string(arquivo))
}
