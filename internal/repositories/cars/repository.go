package cars

import (
	"fmt"
	"github.com/vinicius-robledo/goProjectCRUD/internal/business/model"
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
	r[car.Key] = car
	return car, nil
}

func (r repository) GetAll() ([]model.Car, error) {
	if len(r)==0{
		fmt.Println("Não existem veículos cadastrados")
		//TODO return error OU slice vazia?
	}
	var cars []model.Car
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
	return r[key], nil
}
