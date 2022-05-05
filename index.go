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

const monitoring = 3
const delay = 10

func main() {

	welcome()

	for {
		menuView()
		command := inputUser()
		fmt.Println("")

		switch command {
		case 1:
			initMonitoring()
		case 2:
			fmt.Println("Exibindo Logs...")
			printLogs()
		case 0:
			fmt.Println("Saindo do programa")
			os.Exit(0)
		default:
			fmt.Println("Não conheço esse comando")
			os.Exit(-1)
		}
	}
}

func welcome() {
	name := "Bruno"
	version := 1.1

	fmt.Println("Hello, sr.", name)
	fmt.Println("Version program:", version)
}

func inputUser() int {
	var command int
	fmt.Scan(&command)
	return command
}

func menuView() {
	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do Programa")
}

func initMonitoring() {
	fmt.Println("Monitorando...")
	fmt.Println("")

	pages := readFile()

	for i := 0; i < monitoring; i++ {
		for _, page := range pages {
			verifyPage(page)
			fmt.Println("")
		}
		time.Sleep(delay * time.Second)
		fmt.Println("")
	}

}

func verifyPage(page string) {
	resp, err := http.Get(page)

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	if resp.StatusCode == 200 {
		fmt.Println(page, "foi carregado com sucesso!")
		createLog(page, true)
	} else {
		fmt.Println(page, "não foi carregado status code: ", resp.StatusCode)
		createLog(page, false)
	}
}

func readFile() []string {

	file, err := os.Open("pages.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	var pages []string

	read := bufio.NewReader(file)

	for {
		line, err := read.ReadString('\n')

		if err == io.EOF {
			break

		}
		line = strings.TrimSpace(line)

		pages = append(pages, line)

	}

	file.Close()

	return pages
}

func createLog(page string, status bool) {
	file, err := os.OpenFile("logs.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println(err)
	}

	file.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + page + " - online: " + strconv.FormatBool(status) + "\n")
	file.Close()

}

func printLogs() {
	file, err := ioutil.ReadFile("logs.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro:", file)
	}

	fmt.Println(string(file))

}
