package cars

import (
	"errors"
	"github.com/vinicius-robledo/goProjectCRUD/internal/business/model"
	"github.com/vinicius-robledo/goProjectCRUD/kit"
)

//usar var para variavéis nível 'scope', que estão fora de func
//var repository map[string]model.Car

type repository map[string]model.Car

func CreateCarInterfaceRepository() (InterfaceRepository, error){
	return repository{}, nil
}

	//InterfaceRepository expõe métodos públicos do repository
type InterfaceRepository interface {
	Add(car model.Car) (model.Car, error)
	Update(key string, car model.Car) (model.Car, error)
	Get(key string) (model.Car, error)
	GetAll() ([]model.Car, error)
}

func (r repository) Add(car model.Car) (model.Car, error){
	if car.Key != ""{
		return model.Car{}, errors.New("key não pode ser informado na criação do veículo, será gerada uma chave automaticamente")
	}
	car.Key = kit.GerenateKey()
	r[car.Key] = car
	return car, nil
}

func (r repository) GetAll() ([]model.Car, error) {
	var cars []model.Car
	if len(r)==0{
		return cars, errors.New("não existem veículos cadastrados")
	}

	for _, car :=range r{
		cars = append(cars, car)
	}
	return cars, nil
}

func (r repository) Update(key string, newCar model.Car) (model.Car, error) {
	r[key] = newCar
	return newCar, nil
}

func (r repository) Get(key string) (model.Car, error) {
	car, found := r[key]
	if !found{
		return model.Car{}, errors.New("car not found")
	}
	return car, nil
}
