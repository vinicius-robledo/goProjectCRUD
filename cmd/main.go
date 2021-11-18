package main

import (
	"fmt"
	"github.com/vinicius-robledo/goProjectCRUD/internal/business/car"
	"github.com/vinicius-robledo/goProjectCRUD/internal/business/model"
	"github.com/vinicius-robledo/goProjectCRUD/internal/repositories/cars"
	"os"
)

func main() {

	//kit.RegistraLog("www.teste", false)
	//kit.LeCarrosDoArquivo()

	setupInicial() //criar método de bootstrap
	rep := cars.CreateCarRepository()

	println("Imprimindo type repository:")
	println(rep)

	//iRep := cars.NewCarInterfaceRepository()
	//println("Imprimindo InterfaceRepository:")
	//println(iRep)

	exibeIntroducao()

	for {
		exibeMenu()

		comando := leComando()

		switch comando {
		case 1:
			fmt.Println("Digite dados do veículo...")
			var t, b, y = obterDadosCadastro()
			car := model.New(t,b,y)
			//TODO main chamar SERVICE e service add no Repository
			cars.Add(car.Key, car, rep)
		case 2: //consultar
			fmt.Println("Digite chave do Veículo para CONSULTAR (not implemented)")
			//key := capturarString()
			//car:= repositories.FindById(key, rep)
			//car.PrintCar()
		case 3:
			fmt.Println("Digite a ´KeyCar´ do Veículo para Atualizar: ")
			var k, t, b, y = obterDadosAtualizacao()
			newCar := model.New(t,b,y)
			car.Update(k, newCar, rep)
		case 4:
			fmt.Println("Digite chave do Veículo para DELETAR (not implemented)")

		case 5:
			fmt.Println("Consultando todos os veículos...")
			car.PrintAllCars(rep)
		case 0:
			fmt.Println("Saindo do programa")
			os.Exit(0)
		default:
			fmt.Println("Não conheço este comando")
			os.Exit(-1)
		}
	}

}


func setupInicial() {

}

func obterDadosCadastro() (model string, brand string, year string){
	fmt.Println("Digite a MARCA do Veículo:")
	brand = capturarString()
	fmt.Println("Digite o MODELO do Veículo:")
	model = capturarString()
	fmt.Println("Digite o ANO do Veículo:")
	//year, _ = strconv.Atoi(capturarString())
	year = capturarString()
	return model, brand, year
}


func obterDadosAtualizacao() (key string, model string, brand string, year string) {
	fmt.Println("Digite a KEY do Veículo que quer atualizar:")
	key = capturarString()
	fmt.Println("Digite a MARCA do Veículo:")
	brand = capturarString()
	fmt.Println("Digite o MODELO do Veículo:")
	model = capturarString()
	fmt.Println("Digite o ANO do Veículo:")
	//year, _ = strconv.Atoi(capturarString())
	year = capturarString()
	return key, model, brand, year
}


func exibeIntroducao() {
	versao := 1.0
	fmt.Println("Rodando programa cadastro de Carros. Este programa está na versão", versao)
}

func exibeMenu() {
	fmt.Println(" ")
	fmt.Println("======MENU====== ")
	fmt.Println("1- Cadastrar Veículo")
	fmt.Println("2- Consultar dados do Veículo")
	fmt.Println("3- Atualizar dados de Veículo")
	fmt.Println("4- Deletar Veículo")
	fmt.Println("5- Consultar todos os Veículos")
	fmt.Println("0- Sair do Programa")

}

func leComando() int {
	var comandoLido int
	fmt.Scan(&comandoLido)
	fmt.Println("O comando escolhido foi", comandoLido)

	return comandoLido
}

func capturarString() string {
	var comandoLido string
	fmt.Scan(&comandoLido)
	return comandoLido
}