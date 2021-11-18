package cars

import (
	"fmt"
	"github.com/vinicius-robledo/goProjectCRUD/internal/business/model"
	"strconv"

	//"github.com/vinicius-robledo/goProjectCRUD/internal/repositories"
	//"github.com/vinicius-robledo/goProjectCRUD/kit"
	"testing"
)


func createCars() []model.Car {

	cars:= make([]model.Car, 3)
    cars[0] = model.Car{Title: "M2", Brand: "BMW", Year: "2020"}
	cars[1] = model.Car{Title: "SLK", Brand: "Mercedes", Year: "2005"}
	cars[2] = model.Car{Title: "S3", Brand: "Audi", Year: "2018"}
	return cars

}

func TestCreateCarRepositoryAndAdd(t *testing.T){

	rep := CreateCarRepository()
	cars := createCars()

	fmt.Println("Teste Create Cars. Criando total de " +  strconv.Itoa(len(cars))  + " carros")
	for _, car := range cars{
		//key:= kit.GerenateKey()
		fmt.Printf("%+v", car)
		fmt.Println("")
		Add(car.Key, car, rep)

	}
}

func TestGetAnd(t *testing.T){
	rep := CreateCarRepository()

	oldCar := model.New("M2", "BMW", "2020")
	Add(oldCar.Key, oldCar, rep)

	newCar := Get(oldCar.Key, rep)

	if oldCar != newCar{
		t.Error("Erro ao consultar carro")
	}


}

func TestUpdate(t *testing.T){
	rep := CreateCarRepository()

	oldCar := model.New("M2", "BMW", "2020")

	newCar := model.Car{Key: oldCar.Key, Title: "X6", Brand: "BMW", Year: "2018"}

	Update(newCar.Key, newCar , rep)
}

