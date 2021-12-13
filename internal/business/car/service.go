package car

import (
	"errors"
	"github.com/vinicius-robledo/goProjectCRUD/internal/business/model"
	"github.com/vinicius-robledo/goProjectCRUD/internal/repositories/cars"
)

type service struct{
	rep cars.InterfaceRepository
}

type Service interface {
	CreateCar(c model.Car) (model.Car, error)
	GetCars() ([]model.Car, error)
	Update(keyToUpdate string, newCar model.Car) (model.Car, error)
	GetCarById(key string) (model.Car, error)
	//PrintAllCars()
}

func CreateService(r cars.InterfaceRepository) Service{
	s := service{r}
	return s
}

func (s service) GetCarById(key string) (model.Car, error) {
	// Loop over the list of Car, looking for an car whose ID value matches the parameter.
	listCars, _ := s.GetCars()
	if len(listCars) == 0{
		return model.Car{}, errors.New("não existem carros cadastrados")
	}
	for _, car := range listCars {
		if car.Key == key {
			return car, nil
		}
	}
	return model.Car{}, errors.New("this car Key don't have any reference in database")
}

//func do Service é onde ficam regras de negócio
func (s service)Update(keyToUpdate string, newCar model.Car) (model.Car, error){
	oldCar, err := s.rep.Get(keyToUpdate)

	if err != nil{
		return model.Car{}, errors.New("car not found in repository")
	}

	allRequiredOk, err := validateRequiredField(newCar)

	if !allRequiredOk {
		return model.Car{}, errors.New(err.Error())
	}

	if oldCar.Brand != newCar.Brand{
		println("Não é permitido alterar marca do carro")
		return model.Car{}, errors.New("não é permitido alterar marca do carro. Marca anterior: " + oldCar.Brand  + " | Marca nova: " + newCar.Brand)
	}
	newCar.Key = keyToUpdate
	s.rep.Update(keyToUpdate, newCar)
	return newCar, nil
}

func (s service) GetCars() ([]model.Car, error){
	listCars, _ := s.rep.GetAll()
	if len(listCars) == 0{
		return listCars, errors.New("não existem carros cadastrados")
	}
	return listCars, nil
}

func (s service) CreateCar(c model.Car) (model.Car, error){
	allRequiredOk, err := validateRequiredField(c)
	var newCar model.Car
		if !allRequiredOk {
			return newCar, errors.New(err.Error())
		}else{
			newCar = model.Car{Title: c.Title, Brand: c.Brand, Year: c.Year}
			newCar,_ = s.rep.Add(newCar)
			return newCar, nil
		}

}

func validateRequiredField(c model.Car) (bool, error){
	if c.Key != ""{
		return false, errors.New("key não pode ser informado na criação do veículo, será gerada uma chave automaticamente")
	}else if c.Title == "" {
		return false, errors.New("obrigatório informar o MODELO do veículo")
	}else if c.Brand == ""{
		return false, errors.New("obrigatório informar o MARCA do veículo")
	}else if c.Year == ""{
		return false, errors.New("obrigatório informar o ANO do veículo")
	}else{
		return true, nil
	}
}

// func PrintAllCars usada somente para modo CONSOLE
//func (s service) PrintAllCars(){
//	cars, _ := s.rep.GetAll()
//	if len(cars)==0{
//		fmt.Println("Não existem veículos cadastrados")
//	}
//	for _, car :=range cars{
//		//fmt.Println("Key Car:", key ,"has Properties: { ", car , " }")
//		fmt.Printf("%+v",  car )
//		fmt.Println("")
//	}
//}
