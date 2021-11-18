package cars

import (
	"github.com/vinicius-robledo/goProjectCRUD/internal/business/model"
)

//usar var para variavéis nível 'scope', que estão fora de func
//var repository map[string]model.Car

type Repository map[string]model.Car

func CreateCarRepository() Repository{
	mapCars := Repository{}
	//mapCars := make(map[string] model.Car)
	return mapCars
}

func Add(key string, car model.Car, rep map[string] model.Car){
	rep[key] = car
}


func Update(key string, car model.Car, rep map[string] model.Car){
	rep[key] = car
}

func Get(key string, rep map[string] model.Car) model.Car{
	return rep[key]
}


//func PrintMap(c map[string] model.Car){
//	if len(c)==0{
//		fmt.Println("Não existem veículos cadastrados")
//	}
//	for key, car :=range c{
//		fmt.Println("Key Car:", key ,"has Properties: { ", car , " }")
//		fmt.Printf("%+v",  car )
//	}
//}






//Interface e métodos da interface
//type InterfaceRepository interface {
//	Add(key string, car model.Car, rep InterfaceRepository)
//	AddToInterface(key string, car model.Car, rep InterfaceRepository)
//	Update(key string, car model.Car, rep InterfaceRepository)
//	GetInterface(key string, rep InterfaceRepository) model.Car
//}



//
//func (r repository) AddToInterface(key string, car model.Car, rep InterfaceRepository) {
//	r.AddToInterface(key, car, rep)
//}
//
//func NewCarInterfaceRepository() InterfaceRepository{
//	repository := make(repository)
//	return repository
//}
//
//func AddCarToRep(key string, car model.Car, ir InterfaceRepository) {
//	ir.AddToInterface(key, car, ir)
//}
//
//func (r repository) Add(key string, car model.Car, ir InterfaceRepository) {
//	panic("implement me")
//}
//
//func (r repository) Update(key string, car model.Car, ir InterfaceRepository) {
//	panic("implement me")
//}
//
//func (r repository) GetInterface(key string, ir InterfaceRepository) model.Car {
//	panic("implement me")
//}







//type boot interface{
//	getGreetings() string
//}

// //func que recebe uma interface
//func printGreeting(b boot){
//	fmt.Println(b.getGreetings())
//}
//
//type englishBoot struct {}
//
//func (e englishBoot) getGreetings() string {
//	return "Hello!"
//}
//
//type spanishBoot struct {}
//
//func (s spanishBoot) getGreetings() string {
//	return "Hola!"
//}

//func teste {
//	eb := englishBoot{}
//	sb := spanishBoot{}
//
//	printGreeting(eb)
//	printGreeting(sb)
//
//}


