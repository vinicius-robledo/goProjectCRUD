package api

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/vinicius-robledo/goProjectCRUD/internal/business/car"
	"github.com/vinicius-robledo/goProjectCRUD/internal/business/model"
	"github.com/vinicius-robledo/goProjectCRUD/internal/repositories/cars"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetCarsGin(t *testing.T) {
	//repository := new(cars.RepositoryMock)
	//service := car.CreateService(repository)
	service := car.ServiceMock{}

	type mocks struct {
		GetAllOuput		cars.GetAllOutput
	}
	tt := []struct {
		name        		string
		mock 				mocks
		httpMethod			string
		inputRequest		string
		inputPath			string
		expectedOutput 		string
		expectedErr 		error
		expectedHttpCode	int
	}{
		{
			name: "Success",
			mock: mocks{GetAllOuput:
				cars.GetAllOutput{Cars: []model.Car{{Key: "11c6184d-c848-4848-a7d8-a12e408a4e11", Title: "M2", Brand: "BMW", Year: "2020"},
													{Key: "11c6184d-c848-4848-a7d8-a12e408a4e11", Title: "TT", Brand: "Audi", Year: "2018"}},
									Err: nil}},
			httpMethod: "GET",
			inputPath: "/carsGin",
			inputRequest: "",
			//TODO Input pode ficar no formate STRUCTS ou deixar no JSON/STRING?
			expectedOutput: "[\n    {\n        \"key\": \"11c6184d-c848-4848-a7d8-a12e408a4e11\",\n        \"title\": \"M2\",\n        \"brand\": \"BMW\",\n        \"year\": \"2020\"\n    },\n    {\n        \"key\": \"11c6184d-c848-4848-a7d8-a12e408a4e11\",\n        \"title\": \"TT\",\n        \"brand\": \"Audi\",\n        \"year\": \"2018\"\n    }\n]"   ,
			expectedErr:  nil,
			expectedHttpCode:	http.StatusOK,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			service.On("GetCars").Return(tc.mock.GetAllOuput.Cars , tc.mock.GetAllOuput.Err)
			// Switch to test mode so you don't get such noisy output
			gin.SetMode(gin.TestMode)
			// Setup your router, just like you did in your main function, and register your routes
			router := gin.Default()
			handler := Handler{service: &service}
			//router.GET(tc.inputPath, handler.getCarsGin)
			router.Handle(tc.httpMethod, tc.inputPath, handler.getCarsGin)

			// Create the mock request you'd like to test. Make sure the second argument here is the same as one of the routes you defined in the router setup block!
			req, err := http.NewRequest(tc.httpMethod, tc.inputPath, nil)
			if err != nil {
				t.Fatalf("Couldn't create request: %v\n", err)
			}

			respWriter := httptest.NewRecorder()
			router.ServeHTTP(respWriter, req)

			//UTIL para converter BODY para STRUCT
			//var listCars []model.Car
			//json.NewDecoder(respWriter.Body).Decode(&listCars)

			assert.Equal(t, tc.expectedHttpCode, respWriter.Code)
			assert.Equal(t, tc.expectedOutput, respWriter.Body.String())

		})
	}
}


func TestControllerGetAll(t *testing.T) {
	service := car.ServiceMock{}
	service.On("GetCars").Return([]model.Car{{Key: "22c6184d-c848-4848-a7d8-a12e408a4e22", Title: "C63 AMG", Brand: "Mercedes", Year: "2010"},}, nil )
	// Switch to test mode so you don't get such noisy output
	gin.SetMode(gin.TestMode)
	// Setup your router, just like you did in your main function, and register your routes
	r := gin.Default()
	handler := Handler{service: &service}
	r.GET("/carsGin", handler.getCarsGin)
	//r.GET("/users", GetUsers)

	// Create the mock request you'd like to test. Make sure the second argument here is the same as one of the routes you defined in the router setup block!
	req, err := http.NewRequest(http.MethodGet, "/carsGin", nil)
	if err != nil {
		t.Fatalf("Couldn't create request: %v\n", err)
	}
	// Create a response recorder so you can inspect the response
	respWriter := httptest.NewRecorder()
	// Perform the request
	r.ServeHTTP(respWriter, req)
	// Check to see if the response was what you expected
	if respWriter.Code != http.StatusOK {
		t.Fatalf("Expected to get status %d but instead got %d\n", http.StatusOK, respWriter.Code)
	}

	var listCars []model.Car
	json.NewDecoder(respWriter.Body).Decode(&listCars)


	//fmt.Println(receivedCar)

	// Check to see if the response was what you expected
	if respWriter.Code != http.StatusOK {
		t.Fatalf("Expected to get status %d but instead got %d\n", http.StatusOK, respWriter.Code)
	}

	b, _ := ioutil.ReadAll(respWriter.Body)

	println(string(b))

	//assert.Equal(t, tc.expectedHTTPCode, response.Code)

	//t.Error(w.Code, string(b))
}