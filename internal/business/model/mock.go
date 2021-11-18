package model

import (
	"context"

	"github.com/stretchr/testify/mock"
)

type RepositoryMock struct {
	mock.Mock
}

func (b *RepositoryMock) Add(_ context.Context, car Car) error {
	args := b.Called(car)
	return args.Error(0)
}

//func (b *RepositoryMock) GetByID(_ context.Context, key BatchKey) (Car, error) {
//	args := b.Called(key)
//	return args.Get(0).(Car), args.Error(1)
//}

func (b *RepositoryMock) Update(_ context.Context, car Car) error {
	args := b.Called(car)
	return args.Error(0)
}

