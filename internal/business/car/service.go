package car

import (
	"errors"
	"fmt"
	"github.com/vinicius-robledo/goProjectCRUD/internal/business/model"
	"github.com/vinicius-robledo/goProjectCRUD/internal/repositories/cars"
)

type service struct{
	rep cars.InterfaceRepository
}

type Service interface {
	CreateCar(c model.Car) (model.Car, error)
	GetCars() []model.Car
	Update(keyToUpdate string, newCar model.Car) (model.Car, error)
	GetCarById(key string) (model.Car, error)
}

func CreateService(r cars.InterfaceRepository) Service{
	s := service{r}
	return s
}

func (s service) GetCarById(key string) (model.Car, error) {
	// Loop over the list of Car, looking for an car whose ID value matches the parameter.
	for _, car := range s.GetCars() {
		if car.Key == key {
			return car, nil
		}
	}
	return model.Car{}, errors.New("This car Key don't have any reference in database. ")
}

//func do Service é onde ficam regras de negócio
func (s service)Update(keyToUpdate string, newCar model.Car) (model.Car, error){
	oldCar, err := s.rep.Get(keyToUpdate)

	if err != nil{
		return model.Car{}, errors.New("car not found")
	}

	if oldCar.Brand != newCar.Brand{
		println("Não é permitido alterar marca do carro")
		return model.Car{}, errors.New("não é permitido alterar marca do carro. Marca anterior: " + oldCar.Brand  + " | Marca nova: " + newCar.Brand)
	}
	println("Passou pelo service.Update")
	newCar.Key = keyToUpdate
	s.rep.Update(keyToUpdate, newCar)
	return newCar, nil
}

func (s service) GetCars() []model.Car{
	listCars, _ := s.rep.GetAll()
	return listCars
}

func (s service) CreateCar(c model.Car) (model.Car, error){
	var newCar model.Car
	if c.Title == "" {
		return newCar, errors.New("Obrigatório informar o MODELO do veículo")
	}else if c.Brand == ""{
		return newCar, errors.New("Obrigatório informar o MARCA do veículo")
	}else if c.Year == ""{
		return newCar, errors.New("Obrigatório informar o ANO do veículo")
	}else{
		newCar = model.New(c.Title, c.Brand, c.Year)
		s.rep.Add(newCar)
		return newCar, nil
	}

}

// func PrintAllCars usada somente para modo CONSOLE
func PrintAllCars(c cars.InterfaceRepository){
	cars, _ := c.GetAll()
	if len(cars)==0{
		fmt.Println("Não existem veículos cadastrados")
	}
	for _, car :=range cars{
		//fmt.Println("Key Car:", key ,"has Properties: { ", car , " }")
		fmt.Printf("%+v",  car )
	}
}
