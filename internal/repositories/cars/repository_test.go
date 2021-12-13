package cars

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/vinicius-robledo/goProjectCRUD/internal/business/model"
	"strconv"
	"testing"
)


func createCars() []model.Car {
	cars:= make([]model.Car, 2)
	cars[0] = model.Car{Title: "M2",  Brand: "BMW", Year: "2020"}
	cars[1] = model.Car{Title: "SLK",  Brand: "Mercedes", Year: "2005"}
	//cars[2] = model.Car{Title: "A3",  Brand: "Audi", Year: "2018"}
	return cars
}

func TestCreateCarInterfaceRepository(t *testing.T){
	rep, _ := CreateCarInterfaceRepository()
	cars := createCars()

	fmt.Println("Teste Create Cars. Criando total de " +  strconv.Itoa(len(cars))  + " carros")
	for _, car := range cars{
		//key:= kit.GerenateKey()
		fmt.Printf("%+v", car)
		fmt.Println("")
		rep.Add(car)

	}
}

func TestCreateCarInvalidKey(t *testing.T){
	rep, _ := CreateCarInterfaceRepository()
	car := model.Car{Key: "aaaaaa", Title: "M2", Brand: "BMW", Year: "2015"}
	_, err := rep.Add(car)
	assert.Equal(t, "key não pode ser informado na criação do veículo, será gerada uma chave automaticamente", err.Error())

}

func TestGet(t *testing.T){
	rep, _ := CreateCarInterfaceRepository()

	oldCar := model.Car{"","M2", "BMW", "2020"}
	//oldCar := model.New("M2", "BMW", "2020")
	oldCar, _ = rep.Add(oldCar)

	newCar ,_ := rep.Get(oldCar.Key)

	if oldCar != newCar{
		t.Error("Erro ao consultar carro")
	}
}

func TestGetInvalidKey(t *testing.T){
	rep, _ := CreateCarInterfaceRepository()

	oldCar := model.Car{"","M2", "BMW", "2020"}
	oldCar, _ = rep.Add(oldCar)

	_ ,err := rep.Get("invalid-key")

	assert.Equal(t,"car not found", err.Error())

}

func TestUpdateInterface(t *testing.T){
	rep, _ := CreateCarInterfaceRepository()

	//oldCar := model.New("M2", "BMW", "2020")
	oldCar := model.Car{"","M2", "BMW", "2020"}
	newCar := model.Car{Key: oldCar.Key, Title: "X6", Brand: "BMW", Year: "2018"}

	rep.Update(newCar.Key, newCar)

	validateCar , _ := rep.Get(newCar.Key)
	if validateCar.Title != "X6"{
		t.Error("Erro ao atualizar carro")
	}
}


// GetAll sem TABLE TEST
func TestGetAll(t *testing.T){
	rep, _ := CreateCarInterfaceRepository()
	cars := createCars()

	newCars, _ := rep.GetAll()

	if len(newCars) != 0 {
		t.Error("Erro ao reposutory vazio")
	}

	for i, car :=range cars{
		rep.Add(car)
		println("Adicionando carro " , i, " | Modelo: ", car.Title, " | Marca: ", car.Brand)
	}

	newCars, _ = rep.GetAll()

	for i, car := range newCars{
		println("Consutando carro " , i, " | Modelo: ", car.Title, " | Marca: ", car.Brand)
	}

	if len(newCars) != len(cars) {
		t.Error("Erro ao consultar todos carros")
	}

}


//func TestGetAllTableTest(t *testing.T) {
//	rep, _ := CreateCarInterfaceRepository()
//
//	tt := []struct {
//		name        	string
//		expectOutput 	[]model.Car
//		expectedErr 	error
//	}{
//		{
//			name: "Success",
//			expectOutput: []model.Car{
//				{Key: "11c6184d-c848-4848-a7d8-a12e408a4e11", Title: "M2", Brand: "BMW", Year: "2020"},
//				{Key: "11c6184d-c848-4848-a7d8-a12e408a4e11", Title: "SLK", Brand: "Mercedes", Year: "2010"},
//			},
//			expectedErr:  nil,
//		},
//		{
//			name: "Error_No_Cars",
//			expectOutput: []model.Car{},
//			expectedErr:  errors.New("não existem carros cadastrados"),
//		},
//	}
//
//	for _, tc := range tt {
//		t.Run(tc.name, func(t *testing.T) {
//			cars := createCars()
//			for i, car :=range cars{
//				rep.Add(car)
//				println("Adicionando carro " , i, " | Modelo: ", car.Title, " | Marca: ", car.Brand)
//			}
//			listCars, err := rep.GetAll()
//			require.Equal(t, tc.expectOutput, listCars)
//			require.Equal(t, tc.expectedErr, err)
//		})
//	}
//}



