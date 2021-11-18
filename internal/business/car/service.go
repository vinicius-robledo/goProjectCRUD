package car

import (
	"errors"
	"fmt"
	"github.com/vinicius-robledo/goProjectCRUD/internal/business/model"
	"github.com/vinicius-robledo/goProjectCRUD/internal/repositories/cars"
)


//func do Service é onde ficam regras de negócio
func Update(keytoUpdate string, newCar model.Car,rep cars.Repository) (model.Car, error){
	oldCar := cars.Get(keytoUpdate, rep)
	println("Marca antiga" + oldCar.Brand)
	println("Marca nova" + newCar.Brand)
	if oldCar.Brand != newCar.Brand{
		println("Não é permitido alterar marca do carro")
		return newCar, errors.New("empty name")
	}
	cars.Update(keytoUpdate, newCar, rep)
	return newCar, nil
}

func PrintAllCars(c cars.Repository){
	if len(c)==0{
		fmt.Println("Não existem veículos cadastrados")
	}
	for key, car :=range c{
		fmt.Println("Key Car:", key ,"has Properties: { ", car , " }")
		fmt.Printf("%+v",  car )
	}
}
