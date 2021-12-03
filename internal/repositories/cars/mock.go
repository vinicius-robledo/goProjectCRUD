package cars

import (
	"github.com/stretchr/testify/mock"
	"github.com/vinicius-robledo/goProjectCRUD/internal/business/model"
)

type RepositoryMock struct {
	mock.Mock
}

//func (b *RepositoryMock) Add(_ context.Context, car model.Car) error {
//	args := b.Called(car)
//	return args.Error(0)
//}

func (b *RepositoryMock) Add(car model.Car) (model.Car, error) {
	args := b.Called(car)
	return args.Get(1).(model.Car),args.Error(2)
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