package model

type Car struct {
	Key   string    `json:"key"`
	Title string	`json:"title"`
	Brand string	`json:"brand"`
	Year  string	`json:"year"`
	//PubDate time.Time
	//	Category
}


//TODO remover este metodo do MODEL pois dava diversos problemas nos TESTS
//func New(title string, brand string, year string) Car {
//	key:= kit.GerenateKey()
//	car := Car{Key: key, Title: title, Brand: brand, Year: year}
//	return car
//}
