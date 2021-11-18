package kit

import (
	"bufio"
	"fmt"
	"github.com/google/uuid"
	"io"
	"os"
	"strconv"
	"strings"
	"time"
)

func GerenateKey() string{
	key := fmt.Sprint(uuid.New())
	return key
}


func LeCarrosDoArquivo() []string {

	var sites []string

	arquivo, err := os.Open("cars.txt")
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

	fmt.Println(sites)

	arquivo.Close()

	return sites
}


func registraLog(site string, status bool) {

	arquivo, err := os.OpenFile("log.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site +
		" - online: " + strconv.FormatBool(status) + "\n")

	arquivo.Close()
}