package cars

import (
	"fmt"
	"github.com/vinicius-robledo/goProjectCRUD/internal/business/model"
	"strconv"
	"testing"
)


func createCars() []model.Car {

	cars:= make([]model.Car, 3)
	cars[0] = model.New("M2",  "BMW", "2020")
	cars[1] = model.New("SLK",  "Mercedes", "2005")
	cars[2] = model.New("A3",  "Audi", "2018")
	return cars

}

func TestCreateCarInterfaceRepository(t *testing.T){

	rep, _ := CreateCarInterfaceRepository()
	cars := createCars()

	fmt.Println("Teste Create Cars. Criando total de " +  strconv.Itoa(len(cars))  + " carros")
	for _, car := range cars{
		//key:= kit.GerenateKey()
		fmt.Printf("%+v", car)
		fmt.Println("")
		rep.Add(car)

	}
}

func TestGetInterface(t *testing.T){
	rep, _ := CreateCarInterfaceRepository()

	oldCar := model.New("M2", "BMW", "2020")
	rep.Add(oldCar)

	newCar ,_ := rep.Get(oldCar.Key)

	//newCar := Get(oldCar.Key, rep)

	if oldCar != newCar{
		t.Error("Erro ao consultar carro")
	}

}

func TestUpdateInterface(t *testing.T){
	rep, _ := CreateCarInterfaceRepository()

	oldCar := model.New("M2", "BMW", "2020")
	newCar := model.Car{Key: oldCar.Key, Title: "X6", Brand: "BMW", Year: "2018"}

	rep.Update(newCar.Key, newCar)

	validateCar , _ := rep.Get(newCar.Key)
	if validateCar.Title != "X6"{
		t.Error("Erro ao atualizar carro")
	}
}

func TestGetAll(t *testing.T){
	rep, _ := CreateCarInterfaceRepository()
	cars := createCars()

	for i, car :=range cars{
		rep.Add(car)
		println("Adicionando carro " , i, " | Modelo: ", car.Title, " | Marca: ", car.Brand)
	}

	newCars, _ := rep.GetAll()

	for i, car := range newCars{
		println("Consutando carro " , i, " | Modelo: ", car.Title, " | Marca: ", car.Brand)
	}

	if len(newCars) != len(cars) {
		t.Error("Erro ao consultar todos carros")
	}

}