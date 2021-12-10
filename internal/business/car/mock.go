package car

import (
	"github.com/stretchr/testify/mock"
	"github.com/vinicius-robledo/goProjectCRUD/internal/business/model"
)

type ServiceMock struct {
	mock.Mock
}

//type UpdateInput struct{
//	key string
//	car model.Car
//}
//
//type UpdateOutput struct{
//	car model.Car
//	err error
//}


//type Service interface {
//	CreateCar(c model.Car) (model.Car, error)
//	GetCars() ([]model.Car, error)
//	Update(keyToUpdate string, newCar model.Car) (model.Car, error)
//	GetCarById(key string) (model.Car, error)
//}

func (b *ServiceMock) CreateCar(car model.Car) (model.Car, error) {
	args := b.Called(car)
	return args.Get(0).(model.Car),args.Error(1)
}

func (b *ServiceMock) GetCars() ([]model.Car, error) {
	args := b.Called()
	return args.Get(0).([]model.Car), args.Error(1)
}

func (b *ServiceMock) GetCarById(key string) (model.Car, error) {
	args := b.Called(key)
	return args.Get(0).(model.Car), args.Error(1)
}

func (b *ServiceMock) Update(key string, car model.Car) (model.Car, error) {
	args := b.Called(key, car)
	return args.Get(0).(model.Car), args.Error(1)
}


