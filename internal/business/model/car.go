package model

import "github.com/vinicius-robledo/goProjectCRUD/kit"

type Car struct {
	Key   string    `json:"key"`
	Title string	`json:"title"`
	Brand string	`json:"brand"`
	Year  string	`json:"year"`
	//PubDate time.Time
	//	Category
	//  Owner
}

//TIREI DAQUI e LEVEI para o SERVICE pois dava problemas no MOCK
func New(title string, brand string, year string) Car {
	key:= kit.GerenateKey()
	car := Car{Key: key, Title: title, Brand: brand, Year: year}
	return car
}
