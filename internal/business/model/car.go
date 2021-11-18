package model

import (
	"github.com/vinicius-robledo/goProjectCRUD/kit"
)

type Car struct {
	Key string
	Title string
	Brand string
	Year  string
	//PubDate time.Time
	//	Category
	//  Owner
}

func New(title string, brand string, year string) Car {
	key:= kit.GerenateKey()
	car := Car{Key: key, Title: title, Brand: brand, Year: year}
	return car
}
