package model

import (

	"github.com/stretchr/testify/mock"
)

type ModelMock struct {
	mock.Mock
}

func (b *ModelMock) New(key string, title string, brand string, year string) Car {
	args := b.Called(key, title, brand, year)
	println(args)
	newCar := Car{key, title, brand, year}
	return newCar
}

