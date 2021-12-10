package cars

import (
	"github.com/stretchr/testify/mock"
	"github.com/vinicius-robledo/goProjectCRUD/internal/business/model"
)

type RepositoryMock struct {
	mock.Mock
}

type AddOutput struct{
	Car model.Car
	Err error
}

type GetOutput struct{
	Car model.Car
	Err error
}

type GetAllOutput struct{
	Cars	[]model.Car
	Err 	error
}

type UpdateInput struct{
	Key string
	Car model.Car
}

type UpdateOutput struct{
	Car model.Car
	Err error
}

func (b *RepositoryMock) Add(car model.Car) (model.Car, error) {
	car.Key = "11c6184d-c848-4848-a7d8-a12e408a4e11"
	args := b.Called(car)
	return args.Get(0).(model.Car),args.Error(1)
}

func (b *RepositoryMock) Update(key string, car model.Car) (model.Car, error) {
	args := b.Called(key, car)
	return args.Get(0).(model.Car), args.Error(1)
}

func (b *RepositoryMock) Get(key string) (model.Car, error) {
	args := b.Called(key)
	return args.Get(0).(model.Car), args.Error(1)
}

func (b *RepositoryMock) GetAll() ([]model.Car, error) {
	args := b.Called()
	return args.Get(0).([]model.Car), args.Error(1)
}


//func (b *RepositoryMock) New(key string, title string, brand string, year string) model.Car {
//	args := b.Called(key, title, brand, year)
//	newCar := model.Car{Key: args.Get(0).(string), Title: args.Get(1).(string), Brand: args.Get(2).(string), Year: args.Get(3).(string)}
//	return newCar
//}