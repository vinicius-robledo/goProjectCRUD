package main

import (
	"fmt"
	"github.com/vinicius-robledo/goProjectCRUD/internal/api"
	"github.com/vinicius-robledo/goProjectCRUD/internal/business/car"
	"github.com/vinicius-robledo/goProjectCRUD/internal/repositories/cars"
)

func main() {
	//kit.RegistraLog("www.teste", false)
	//kit.LeCarrosDoArquivo()

	setupInicial() //criar método de bootstrap
	interfaceRep, _ := cars.CreateCarInterfaceRepository()
	exibeIntroducao()

	//iniciar Cars Service e passar para o API o service?
	service := car.CreateService(interfaceRep)

	//TMP adc carros para testar GET
	//car1 := model.New("M2",  "BMW", "2020")
	//car2 := model.New("TT",  "Audi", "2018")
	//car1 := model.Car{"","M2",  "BMW", "2020"}
	//car2 := model.Car{"","TT",  "Audi", "2018"}
	//service.CreateCar(car1)
	//service.CreateCar(car2)

	api.InitHttpServer(service)

	//InitCommandLine(service)

}

// func InitCommandLine utilizada somente para testes no CONSOLE ao invés de WebApp
//func InitCommandLine(s car.Service) {
//	for {
//		exibeMenu()
//
//		comando := leComando()
//
//		switch comando {
//		case 1:
//			fmt.Println("Digite dados do veículo...")
//			var t, b, y = obterDadosCadastro()
//			car := model.Car{"",t,b,y}
//			s.CreateCar(car)
//		case 2: //consultar
//			fmt.Println("Digite chave do Veículo para CONSULTAR (not implemented)")
//			//key := capturarString()
//			//car:= repositories.FindById(key, rep)
//			//car.PrintCar()
//		case 3:
//			fmt.Println("Digite a ´KeyCar´ do Veículo para Atualizar: ")
//			//var k, t, b, y = obterDadosAtualizacao()
//			//newCar := model.New(t,b,y)
//			//car.Update(k, newCar, c)
//			//interfaceRep.UpdateInterface(k, newCar)
//		case 4:
//			fmt.Println("Digite chave do Veículo para DELETAR (not implemented)")
//		case 5:
//			fmt.Println("Consultando todos os veículos...")
//			s.PrintAllCars()
//		case 0:
//			fmt.Println("Saindo do programa")
//			os.Exit(0)
//		default:
//			fmt.Println("Não conheço este comando")
//			os.Exit(-1)
//		}
//	}
//}


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