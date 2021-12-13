package car

import (
	"errors"
	"github.com/stretchr/testify/require"
	"github.com/vinicius-robledo/goProjectCRUD/internal/business/model"
	"github.com/vinicius-robledo/goProjectCRUD/internal/repositories/cars"
	"testing"
)


func TestCreateCar(t *testing.T) {
	repository := new(cars.RepositoryMock)
	service := CreateService(repository)

	type mocks struct {
		AddInput		model.Car
		AddOuput		cars.AddOutput
		//NewId			string
		//err           error
	}
	tt := []struct {
		name        	string
		mock 			mocks
		input       	model.Car
		expectOutput 	model.Car
		expectedErr 	error
	}{

		{
			name:  "Success",
			mock: mocks{AddInput:	model.Car{Key: "11c6184d-c848-4848-a7d8-a12e408a4e11", Title: "ML 500", Brand: "Mercedes", Year: "2010"},
						AddOuput:	cars.AddOutput{ Car: model.Car{Key: "11c6184d-c848-4848-a7d8-a12e408a4e11", Title: "ML 500", Brand: "Mercedes", Year: "2010"} , Err: nil},
			},
			input: model.Car{Title: "ML 500", Brand: "Mercedes", Year: "2010"},
			expectOutput: model.Car{Key: "11c6184d-c848-4848-a7d8-a12e408a4e11", Title: "ML 500", Brand: "Mercedes", Year: "2010"},
			expectedErr: nil,
		},
		{
			name:  "Erro_not_empty_key",
			mock: mocks{AddInput:	model.Car{Key: "11c6184d-c848-4848-a7d8-a12e408a4e11", Title: "ML 500", Brand: "Mercedes", Year: "2010"},
						AddOuput:	cars.AddOutput{ Car: model.Car{Key: "11c6184d-c848-4848-a7d8-a12e408a4e11", Title: "ML 500", Brand: "Mercedes", Year: "2010"} , Err: nil},
			},
			input: model.Car{Key: "11c6184d-c848-4848-a7d8-a12e408a4e11", Title: "ML 500", Brand: "Mercedes", Year: "2010"},
			expectedErr: errors.New("key não pode ser informado na criação do veículo, será gerada uma chave automaticamente"),
		},
		{
			name:  "Error_Empty_Title",
			mock: mocks{AddInput:	model.Car{Title: "", Brand: "Mercedes", Year: "2010"},
						AddOuput:	cars.AddOutput{ Car: model.Car{Key: "11c6184d-c848-4848-a7d8-a12e408a4e11", Title: "ML 500", Brand: "Mercedes", Year: "2010"} , Err: nil},
			},
			input: model.Car{Title: "", Brand: "Mercedes", Year: "2010"},
			expectedErr: errors.New("obrigatório informar o MODELO do veículo"),
		},
		{
			name:  "Error_Empty_Brand",
			mock: mocks{AddInput:	model.Car{Title: "ML 500", Brand: "", Year: "2010"},
						AddOuput:	cars.AddOutput{ Car: model.Car{Key: "11c6184d-c848-4848-a7d8-a12e408a4e11", Title: "ML 500", Brand: "Mercedes", Year: "2010"} , Err: nil},
			},
			input: model.Car{Title: "ML 500", Brand: "", Year: "2010"},
			expectedErr: errors.New("obrigatório informar o MARCA do veículo"),
		},
		{
			name:  "Error_Empty_Year",
			mock: mocks{AddInput:	model.Car{Title: "ML 500", Brand: "Mercedes", Year: ""},
						AddOuput:	cars.AddOutput{ Car: model.Car{Key: "11c6184d-c848-4848-a7d8-a12e408a4e11", Title: "ML 500", Brand: "Mercedes", Year: "2010"} , Err: nil},
			},
			input: model.Car{Title: "ML 500", Brand: "Mercedes", Year: ""},
			expectedErr: errors.New("obrigatório informar o ANO do veículo"),
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			repository.On("Add", tc.mock.AddInput).Return(tc.mock.AddOuput.Car ,tc.mock.AddOuput.Err).Once()
			car, err := service.CreateCar(tc.input)
			require.Equal(t, tc.expectOutput, car)
			require.Equal(t, tc.expectedErr, err)
			//time.Sleep(2 * time.Second)
		})
	}
}

func TestGetCars(t *testing.T) {
	repository := new(cars.RepositoryMock)
	service := CreateService(repository)

	type mocks struct {
		GetAllOuput		cars.GetAllOutput
	}
	tt := []struct {
		name        	string
		mock 			mocks
		expectOutput 	[]model.Car
		expectedErr 	error
	}{
		{
			name: "Success",
			mock: mocks{GetAllOuput: cars.GetAllOutput{Cars: []model.Car{
				{Key: "11c6184d-c848-4848-a7d8-a12e408a4e11", Title: "ML 500", Brand: "Mercedes", Year: "2010"},
				{Key: "11c6184d-c848-4848-a7d8-a12e408a4e11", Title: "M2", Brand: "BMW", Year: "2020"}},
				Err: nil}},
			expectOutput: []model.Car{
				{Key: "11c6184d-c848-4848-a7d8-a12e408a4e11", Title: "ML 500", Brand: "Mercedes", Year: "2010"},
				{Key: "11c6184d-c848-4848-a7d8-a12e408a4e11", Title: "M2", Brand: "BMW", Year: "2020"}},
			expectedErr:  nil,
		},
		{
			name: "Error_No_Cars",
			mock: mocks{GetAllOuput: cars.GetAllOutput{Cars: []model.Car{},
				Err: errors.New("não existem veículos cadastrados")}},
			expectOutput: []model.Car{},
			expectedErr:  errors.New("não existem carros cadastrados"),
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			repository.On("GetAll").Return(tc.mock.GetAllOuput.Cars ,tc.mock.GetAllOuput.Err).Once()

			listCars, err := service.GetCars()
			require.Equal(t, tc.expectOutput, listCars)
			require.Equal(t, tc.expectedErr, err)
		})
	}
}

func TestGetCarById(t *testing.T) {
	repository := new(cars.RepositoryMock)
	service := CreateService(repository)

	type mocks struct {
		GetInput		string
		GetOutput		cars.GetOutput
		GetAllOuput		cars.GetAllOutput
	}
	tt := []struct {
		name        	string
		Input			string
		mock 			mocks
		expectOutput 	model.Car
		expectedErr 	error
	}{
		{
			name: "Success",
			Input: "11c6184d-c848-4848-a7d8-a12e408a4e11",
			mock: mocks{GetInput: "11c6184d-c848-4848-a7d8-a12e408a4e11",
				GetOutput: cars.GetOutput{Car: model.Car{Key: "11c6184d-c848-4848-a7d8-a12e408a4e11", Title: "ML 500", Brand: "Mercedes", Year: "2010"},
					Err: nil},
				GetAllOuput: cars.GetAllOutput{Cars: []model.Car{
					{Key: "11c6184d-c848-4848-a7d8-a12e408a4e11", Title: "ML 500", Brand: "Mercedes", Year: "2010"},
					{Key: "22c6184d-c848-4848-a7d8-a12e408a4e22", Title: "M2", Brand: "BMW", Year: "2020"}},
				},
			},
			expectOutput: model.Car{Key: "11c6184d-c848-4848-a7d8-a12e408a4e11", Title: "ML 500", Brand: "Mercedes", Year: "2010"},
			expectedErr:  nil,
		},
		{
			name: "Error_Zero_Cars",
			Input: "00c6184d-c848-4848-a7d8-a12e408a4e00",
			mock: mocks{GetInput: "00c6184d-c848-4848-a7d8-a12e408a4e00",
				GetOutput: cars.GetOutput{ Car: model.Car{},
					Err: nil},
			},
			expectOutput: model.Car{},
			expectedErr:  errors.New("não existem carros cadastrados"),
		},
		{
			name: "Error_Key_Car_Not_Found",
			Input: "99c6184d-c848-4848-a7d8-a12e408a4e99",
			mock: mocks{GetInput: "99c6184d-c848-4848-a7d8-a12e408a4e99",
				GetOutput: cars.GetOutput{ Car: model.Car{},
					Err: errors.New("car not found")},
				GetAllOuput: cars.GetAllOutput{Cars: []model.Car{
					{Key: "11c6184d-c848-4848-a7d8-a12e408a4e11", Title: "ML 500", Brand: "Mercedes", Year: "2010"},
					{Key: "22c6184d-c848-4848-a7d8-a12e408a4e22", Title: "M2", Brand: "BMW", Year: "2020"}},
				},
			},
			expectOutput: model.Car{},
			expectedErr:  errors.New("this car Key don't have any reference in database"),
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			repository.On("Get", tc.Input).Return(tc.mock.GetOutput.Car, tc.mock.GetOutput.Err).Once()
			repository.On("GetAll").Return(tc.mock.GetAllOuput.Cars, tc.mock.GetAllOuput.Err).Once()

			listCars, err := service.GetCarById(tc.Input)
			require.Equal(t, tc.expectOutput, listCars)
			require.Equal(t, tc.expectedErr, err)
		})
	}
}


func TestUpdate(t *testing.T) {
	repository := new(cars.RepositoryMock)
	service := CreateService(repository)

	type mocks struct {
		//carMockRepository	*cars.RepositoryMock
		GetInput		string
		GetOuput		cars.GetOutput
		UpdateInput		cars.UpdateInput
		UpdateOutput	cars.UpdateOutput
		//err               	error
	}
	tt := []struct {
		name        string
		mock 		mocks
		input        []model.Car
		expectOutput model.Car
		expectedErr error
	}{
		{
			name:  "Success",
			mock: mocks{GetInput: "11c6184d-c848-4848-a7d8-a12e408a4e11",
						GetOuput:		cars.GetOutput{ Car: model.Car{Key: "11c6184d-c848-4848-a7d8-a12e408a4e11", Title: "ML 500", Brand: "Mercedes", Year: "2010"} , Err: nil},
						UpdateInput:	cars.UpdateInput{ Key: "11c6184d-c848-4848-a7d8-a12e408a4e11",  Car: model.Car{Key: "11c6184d-c848-4848-a7d8-a12e408a4e11", Title: "SLK", Brand: "Mercedes", Year: "2005"}},
						UpdateOutput:	cars.UpdateOutput{Car: model.Car{Key: "11c6184d-c848-4848-a7d8-a12e408a4e11", Title: "SLK", Brand: "Mercedes", Year: "2005"}, Err: nil},
						//err: nil,
			},
			input: []model.Car{ {Key: "11c6184d-c848-4848-a7d8-a12e408a4e11", Title: "ML 500", Brand: "Mercedes", Year: "2010"},
								{Title: "SLK", Brand: "Mercedes", Year: "2005"}},
			expectOutput: model.Car{Key: "11c6184d-c848-4848-a7d8-a12e408a4e11", Title: "SLK", Brand: "Mercedes", Year: "2005"},
			expectedErr: nil,
		},
		{
			name:  "Error_Another_Brand",
			mock: mocks{GetInput: "11c6184d-c848-4848-a7d8-a12e408a4e11",
						GetOuput:    cars.GetOutput{ Car: model.Car{Key: "11c6184d-c848-4848-a7d8-a12e408a4e11", Title: "ML 500", Brand: "Mercedes", Year: "2010"} , Err: nil},
						UpdateInput:  cars.UpdateInput{ Key: "11c6184d-c848-4848-a7d8-a12e408a4e11",  Car: model.Car{Key: "11c6184d-c848-4848-a7d8-a12e408a4e11", Title: "Mercedes", Brand: "BMW", Year: "2008"}},
						UpdateOutput: cars.UpdateOutput{Car: model.Car{Key: "11c6184d-c848-4848-a7d8-a12e408a4e11", Title: "ML 500", Brand: "Mercedes", Year: "2010"}, Err: nil},
				//err: nil,
			},
			input: []model.Car{ {Key: "11c6184d-c848-4848-a7d8-a12e408a4e11", Title: "ML 500", Brand: "Mercedes", Year: "2010"},
								{Title: "M2", Brand: "BMW", Year: "2008"}},
			expectOutput: model.Car{},
			expectedErr: errors.New("não é permitido alterar marca do carro. Marca anterior: "  +  "Mercedes" +  " | Marca nova: " + "BMW"),
		},
		{
			name:  "Error_Old_Car_Not_Found",
			mock: mocks{GetInput: "99c6184d-c848-4848-a7d8-a12e408a4e99",
						GetOuput:    cars.GetOutput{ Car: model.Car{} , Err: errors.New("car not found")},
			},
			input: []model.Car{ {Key: "99c6184d-c848-4848-a7d8-a12e408a4e99", Title: "ML 500", Brand: "Mercedes", Year: "2010"},
								{Title: "SLK", Brand: "Mercedes", Year: "2005"}},
			expectOutput: model.Car{},
			expectedErr: errors.New("car not found in repository"),
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			repository.On("Get", tc.mock.GetInput).Return(tc.mock.GetOuput.Car ,tc.mock.GetOuput.Err)
			repository.On("Update",tc.mock.UpdateInput.Key, tc.mock.UpdateInput.Car).Return(tc.mock.UpdateOutput.Car ,tc.mock.UpdateOutput.Err)

			car, err := service.Update(tc.input[0].Key, tc.input[1])
			require.Equal(t, tc.expectOutput, car)
			require.Equal(t, tc.expectedErr, err)
		})
	}
}


//*** TESTs sem usar padrao TestTable

// *** SEM MOCK ***
//var interfaceRep, _ = cars.CreateCarInterfaceRepository()
//var s = CreateService(interfaceRep)
//
//func TestUpdateSemMock(t *testing.T) {
//	newCar := model.New("M2", "BMW", "2020")
//
//	if newCar.Key == "" || len(newCar.Key) != 36{
//		t.Error("func New Car não gerou Key com 36 posições para novo veículo")
//	}
//
//	s.Update(newCar.Key, newCar)
//	assert.Equal(t, newCar.Title, "M2", "Erro no MODELO do carro")
//	assert.Equal(t, newCar.Brand, "BMW", "Erro na MARCA do carro")
//
//}

// *** COM MOCK ***
//func TestUpdate(t *testing.T) {
//	testObjMock := new(cars.RepositoryMock)
//
//
//	oldCar := model.Car{Key: "11c6184d-c848-4848-a7d8-a12e408a4e11", Title: "ML 500", Brand: "Mercedes", Year: "2010"}
//	newCarSameBrand := model.Car{Key: "11c6184d-c848-4848-a7d8-a12e408a4e11", Title: "SLK", Brand: "Mercedes", Year: "2015"}
//	newCarAnotherBrand := model.Car{Key: "11c6184d-c848-4848-a7d8-a12e408a4e11", Title: "M2", Brand: "BMW", Year: "2006"}
//	newCarIdNotFound := model.Car{Key: "99c6184d-c848-4848-a7d8-a12e408a4e99", Title: "M2", Brand: "BMW", Year: "2006"}
//	// setup MOCK expectations
//	testObjMock.On("Update",oldCar.Key, newCarSameBrand).Return(newCarSameBrand ,nil)
//	testObjMock.On("Update",oldCar.Key, newCarAnotherBrand).Return(newCarAnotherBrand ,nil)
//	testObjMock.On("Get", "11c6184d-c848-4848-a7d8-a12e408a4e11").Return(oldCar ,nil)
//	testObjMock.On("Get", newCarIdNotFound.Key).Return(model.Car{} , errors.New("car not found"))
//
//	service := CreateService(testObjMock)
//	newCarSameBrand, err1 := service.Update(oldCar.Key, newCarSameBrand)
//	_ , err2 := service.Update(oldCar.Key, newCarAnotherBrand)
//	_ , err3 := service.Update(newCarIdNotFound.Key, newCarIdNotFound)
//
//	if err1 != nil {
//		t.Error("Erro ao atualizar carro. | Error: ", err1)
//	}
//	assert.Equal(t, newCarSameBrand.Title, "SLK", "Erro no MODELO do carro")
//
//
//	if err2 == nil {
//		t.Error("permitiu alteração de Marca do Veículo. | Error: ", err2)
//	}
//	assert.NotEqual(t, err2, nil, "Erro no update do MODELO do carro")
//	assert.Equal(t, err2, errors.New("não é permitido alterar marca do carro. Marca anterior: " + oldCar.Brand  + " | Marca nova: " + newCarAnotherBrand.Brand)	, "Erro no update do MODELO do carro")
//
//	if err3 == nil {
//		t.Error("permitiu alteração de carro  não existente. | Error: ", err2)
//	}
//	assert.NotEqual(t, err3, nil, "permitiu alteração de carro  não existente")
//}


