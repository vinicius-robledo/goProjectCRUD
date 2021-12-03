package car

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/vinicius-robledo/goProjectCRUD/internal/business/model"
	"github.com/vinicius-robledo/goProjectCRUD/internal/repositories/cars"
	"testing"
)

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
func TestUpdate(t *testing.T) {
	// create an instance of our test object
	testObjMock := new(cars.RepositoryMock)

	// setup expectations
	oldCar := model.Car{Key: "11c6184d-c848-4848-a7d8-a12e408a4e11", Title: "ML 500", Brand: "Mercedez", Year: "2010"}
	newCarSameBrand := model.Car{Key: "11c6184d-c848-4848-a7d8-a12e408a4e11", Title: "SLK", Brand: "Mercedez", Year: "2015"}
	newCarAnotherBrand := model.Car{Key: "11c6184d-c848-4848-a7d8-a12e408a4e11", Title: "M2", Brand: "BMW", Year: "2006"}


	testObjMock.On("Update",oldCar.Key, newCarSameBrand).Return(newCarSameBrand ,nil)
	testObjMock.On("Update",oldCar.Key, newCarAnotherBrand).Return(newCarAnotherBrand ,nil)

	testObjMock.On("Get", "11c6184d-c848-4848-a7d8-a12e408a4e11").Return(oldCar ,nil)

	//iniciar Cars Service
	service := CreateService(testObjMock)
	newCarSameBrand, err1 := service.Update(oldCar.Key, newCarSameBrand)
	newCarAnotherBrand, err2 := service.Update(oldCar.Key, newCarAnotherBrand)

	if err1 != nil {
		t.Error("Erro ao atualizar carro. | Error: ", err1)
	}

	if err2 == nil {
		t.Error("permitiu alteração de Marca do Veículo. | Error: ", err2)
	}

	// assert that the expectations were met
	//testObjMock.AssertExpectations(t)

	assert.Equal(t, newCarSameBrand.Title, "SLK", "Erro no MODELO do carro")
	assert.NotEqual(t, newCarAnotherBrand.Title, oldCar.Title, "Erro no MODELO do carro")

}


//func TestUpdateTableTest(t *testing.T) {
//	repo, _ := cars.CreateCarInterfaceRepository()
//
//	tt := []struct {
//		name        string
//		input       model.Car
//		expectedErr error
//	}{
//		{
//			name:  "Success",
//			input: mockCar("success"),
//			expectedErr: nil,
//		},
//		{
//			name:  "Error_Batch_Number_Validation",
//			input: mockCar("error_batch_number_validation"),
//			expectedErr: errors.New("error validating required batch data for key"),
//		},
//		//{
//		//	name:  "Error_Dynamo_Conditional",
//		//	input: mockCar("error_dynamo_conditional"),
//		//	expectedErr: model.BuildErrConditionalError(map[string]string{"batch_id": "error_dynamo_conditional", "batch_number": "987654"}),
//		//},
//		//{
//		//	name:  "Error_Internal",
//		//	input: mockCar("error_internal"),
//		//	expectedErr: errors.New("error adding batch: InternalServerError: internal error"),
//		//},
//		//{
//		//	name:  "Error_Unexpected",
//		//	input: mockCar("error_unexpected"),
//		//	expectedErr: errors.New("error adding batch: unexpected error"),
//		//},
//	}
//
//	for _, tc := range tt {
//		t.Run(tc.name, func(t *testing.T) {
//			//dynamoMock.On("PutItem", context.Background(), tc.mockData.in, mock.Anything).Return(tc.mockData.out, tc.mockData.err)
//			err := repo.Add(context.Background(), tc.input)
//			require.Equal(t, tc.expectedErr, err)
//		})
//	}
//}


//TODO validar teste. Problema com o KeyCar randomico
//func TestCreateCar(t *testing.T) {
//
//	testRepositoryMock := new(cars.RepositoryMock)
//
//	// setup expectations
//	//receivedCar := testRepositoryMock.New("99c6184d-c848-4848-a7d8-a12e408a4e79", "SLK", "Mercedez",  "2015")
//	receivedCar := model.Car{Key: "99c6184d-c848-4848-a7d8-a12e408a4e79", Title: "SLK", Brand: "Mercedez", Year: "2015"}
//
//	testRepositoryMock.On("Add", receivedCar).Return(receivedCar, nil)
//	//testRepositoryMock.On("New", receivedCar.Key, receivedCar.Title, receivedCar.Brand, receivedCar.Year).Return(receivedCar, nil)
//	//testModelMock.On("New", receivedCar.Key ,receivedCar.Title, receivedCar.Brand, receivedCar.Year).Return(receivedCar)
//
//	//iniciar Cars Service
//	service := CreateService(testRepositoryMock)
//	service.CreateCar(receivedCar)
//
//	// call the code we are testing
//	//targetFuncThatDoesSomethingWithObjMock(testObjMock, newCar)
//
//	// assert that the expectations were met
//	testRepositoryMock.AssertExpectations(t)
//
//}

func targetFuncThatDoesSomethingWithObjMock(objMock *cars.RepositoryMock, c model.Car) {
	objMock.Add(c)
}




/*
  Test objects
*/

// MyMockedObject is a mocked object that implements an interface
// that describes an object that the code I am testing relies on.
type MyMockedObject struct{
	mock.Mock
}

// DoSomething is a method on MyMockedObject that implements some interface and just records the activity,
//and returns what the Mock object tells it to.
// In the real object, this method would do something useful, but since this is a mocked object - we're just going to stub it out.
//
// NOTE: This method is not being tested here, code that uses this object is.
func (m *MyMockedObject) DoSomething(number int) (bool, error) {

	args := m.Called(number)
	return args.Bool(0), args.Error(1)

}

/*
  Actual test functions
*/

// TestSomething is an example of how to use our test object to
// make assertions about some target code we are testing.
func TestSomething(t *testing.T) {

	// create an instance of our test object
	testObj := new(MyMockedObject)

	// setup expectations
	testObj.On("DoSomething", 123).Return(true, nil)

	// call the code we are testing
	targetFuncThatDoesSomethingWithObj(testObj)

	// assert that the expectations were met
	testObj.AssertExpectations(t)
}



func targetFuncThatDoesSomethingWithObj(obj *MyMockedObject) {
	obj.DoSomething(123)
	fmt.Println("passou targetFuncThatDoesSomethingWithObj")
}




func mockCar(carID string) model.Car {
	return model.Car{
		Key:           "abcd",
		Title:       "Fusca",
		Brand: "Volkswagem",
		Year:    "1975",
	}
}