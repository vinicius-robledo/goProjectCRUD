package api

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/vinicius-robledo/goProjectCRUD/internal/business/model"

	//"github.com/mercadolibre/fury_go-core/pkg/web"
	"github.com/vinicius-robledo/goProjectCRUD/internal/business/car"
	"net/http"
)

type customHandler struct{}

//Handler joins components to attend requisitions.
type Handler struct {
	service                     car.Service
}

func InitHttpServer(s car.Service) (){
	router := gin.Default()
	handler := Handler{service: s}
	router.GET("/carsGin", handler.getCarsGin)
	router.POST("/carsGin", handler.postCarsGin)
	router.GET("/carsGin/:key", handler.getCarByIDGin)
	router.PATCH("/carsGin/:key", handler.updateCarGin)

	router.Run("localhost:8081")
}


// getCarsGin devolve a lista de Carros em JSON usando gin.Context
func (h Handler) getCarsGin(c *gin.Context) {
	listCars, err := h.service.GetCars()
	if err != nil {
		c.IndentedJSON(http.StatusUnprocessableEntity, err.Error())
	}else{
		c.IndentedJSON(http.StatusOK, listCars)
		//c.JSON(http.StatusOK, gin.H{"data": listCars})
	}



}

// postCarsGin recebe um carro em JSON usando gin.Context e salva no repo
func (h Handler) postCarsGin(c *gin.Context) {
	var receivedCar model.Car
	if err := c.BindJSON(&receivedCar); err != nil {
		c.IndentedJSON(http.StatusUnprocessableEntity, errors.New("this body-request format isn't recognized like a car"))
		return
	}
	newCar, err := h.service.CreateCar(receivedCar)

	if err == nil{
		c.IndentedJSON(http.StatusCreated, newCar)
	}else{
		c.IndentedJSON(http.StatusUnprocessableEntity, err.Error())
	}

}

// getCarByIDGin locates the car whose ID value matches the id
func (h Handler) getCarByIDGin(c *gin.Context) {
	keyCar := c.Param("key")

	car, err := h.service.GetCarById(keyCar)

	if err == nil{
		c.IndentedJSON(http.StatusOK, car)
	}else{
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "car not found"})
	}

}

func (h Handler) updateCarGin(c *gin.Context) {
	key := c.Param("key")

	var receivedCar model.Car

	if err := c.BindJSON(&receivedCar); err != nil {
		c.String(http.StatusUnprocessableEntity, "this body-request format isn't recognized like a car")
		return
	}
	newCar, err :=  h.service.Update(key, receivedCar)

	if err == nil{
		c.IndentedJSON(http.StatusOK, newCar)
	}else{
		c.String(http.StatusUnprocessableEntity, err.Error())
	}

}

//// getCarsJSON devolve a lista de Carros em JSON
//func getCarsJSON(w http.ResponseWriter, r *http.Request) {
//	//TODO receber instancia do Service ao inves de criar
//	//car.CreateService().GetCars()
//	listCars := car.Service.GetCars
//	println(listCars)
//	//c.IndentedJSON(http.StatusOK, listCars)
//}




//func (h Handler) getClearingOrderByID(w http.ResponseWriter, r *http.Request) error {
//	clearingOrderID, err := web.Params(r).String("clearingOrderID")
//	if err != nil {
//		return handleError(r.Context(), w, err)
//	}
//	ctx := log.With(r.Context(), log.String("clearingOrderID", clearingOrderID))
//	clearingOrder, err := h.queryService.GetClearingOrderByID(ctx, clearingOrderID)
//	if err != nil {
//	return handleError(ctx, w, err)
//	}
//
//	return web.RespondJSON(w, clearingOrder, http.StatusOK)
//}
