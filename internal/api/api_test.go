package api

import (
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/vinicius-robledo/goProjectCRUD/internal/business/car"
	"github.com/vinicius-robledo/goProjectCRUD/internal/business/model"
	"github.com/vinicius-robledo/goProjectCRUD/internal/repositories/cars"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGetCarsGin(t *testing.T) {
	service := car.ServiceMock{}
	handlerInit := Handler{service: &service}

	type mocks struct {
		GetAllOuput		cars.GetAllOutput
		GetOuput		cars.GetOutput
	}
	tt := []struct {
		name        		string
		mock 				mocks
		httpMethod			string
		inputPath			string
		inputParam			string
		expectedOutput 		string
		expectedHttpCode	int
	}{
		{
			name: "Get_All_Cars_GIN_Success",
			mock: mocks{GetAllOuput:
						cars.GetAllOutput{Cars: []model.Car{{Key: "11c6184d-c848-4848-a7d8-a12e408a4e11", Title: "M2", Brand: "BMW", Year: "2020"},
															{Key: "22c6184d-c848-4848-a7d8-a12e408a4e22", Title: "TT", Brand: "Audi", Year: "2018"}},
						Err: nil}},
			httpMethod: "GET",
			inputPath: "/carsGin",
			//TODO Input pode ficar no formate STRUCTS ou deixar no JSON/STRING?
			expectedOutput: "[\n    {\n        \"key\": \"11c6184d-c848-4848-a7d8-a12e408a4e11\",\n        \"title\": \"M2\",\n        \"brand\": \"BMW\",\n        \"year\": \"2020\"\n    },\n    {\n        \"key\": \"22c6184d-c848-4848-a7d8-a12e408a4e22\",\n        \"title\": \"TT\",\n        \"brand\": \"Audi\",\n        \"year\": \"2018\"\n    }\n]"   ,
			expectedHttpCode:	http.StatusOK,
		},
		{
			name: "Get_All_Cars_GIN_Error_Zero_Cars",
			mock: mocks{GetAllOuput:
						cars.GetAllOutput{Cars: []model.Car{{}},
							Err: errors.New("não existem carros cadastrados")}},
			httpMethod: "GET",
			inputPath: "/carsGin/",
			//inputRequest: "",
			expectedOutput: "\"não existem carros cadastrados\"" ,
			expectedHttpCode:	http.StatusUnprocessableEntity,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			service.On("GetCars").Return(tc.mock.GetAllOuput.Cars , tc.mock.GetAllOuput.Err).Once()
			gin.SetMode(gin.TestMode)
			router := gin.Default()
			//handler := Handler{service: &service}
			router.Handle(tc.httpMethod, tc.inputPath, handlerInit.getCarsGin)

			// Create the mock request you'd like to test. Make sure the second argument here is the same as one of the routes you defined in the router setup block!
			req, err := http.NewRequest(tc.httpMethod, tc.inputPath, nil)
			if err != nil {
				t.Fatalf("Couldn't create request: %v\n", err)
			}

			respWriter := httptest.NewRecorder()
			router.ServeHTTP(respWriter, req)

			assert.Equal(t, tc.expectedHttpCode, respWriter.Code)
			assert.Equal(t, tc.expectedOutput, respWriter.Body.String())

		})
	}
}


func TestGetByIdCarsGin(t *testing.T) {
	service := car.ServiceMock{}
	handlerInit := Handler{service: &service}

	type mocks struct {
		GetAllOuput		cars.GetAllOutput
		GetOuput		cars.GetOutput
	}
	tt := []struct {
		name        		string
		mock 				mocks
		httpMethod			string
		inputPath			string
		inputParamName      string
		inputParamValue		string
		expectedOutput 		string
		expectedHttpCode	int
	}{
		{
			name: "Get_byId_Cars_GIN_Success",
			mock: mocks{GetOuput: cars.GetOutput{Car: model.Car{Key: "11c6184d-c848-4848-a7d8-a12e408a4e11", Title: "M2", Brand: "BMW", Year: "2020"} ,
				Err: nil}},
			httpMethod: "GET",
			inputPath: "/carsGin/",
			inputParamName: ":key",
			inputParamValue: "11c6184d-c848-4848-a7d8-a12e408a4e11",
			expectedOutput: "{\n    \"key\": \"11c6184d-c848-4848-a7d8-a12e408a4e11\",\n    \"title\": \"M2\",\n    \"brand\": \"BMW\",\n    \"year\": \"2020\"\n}"  ,
			expectedHttpCode:	http.StatusOK,
		},
		{
			name: "Get_byId_Cars_Id_NotFound",
			mock: mocks{GetOuput: cars.GetOutput{Car: model.Car{} ,
				Err: errors.New("this car Key don't have any reference in database")}},
			httpMethod: "GET",
			inputPath: "/carsGin/",
			inputParamName: ":key",
			inputParamValue: "99c6184d-c848-4848-a7d8-a12e408a4e99",
			expectedOutput: "{\n    \"message\": \"car not found\"\n}"  ,
			expectedHttpCode:	http.StatusNotFound,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			service.On("GetCarById", tc.inputParamValue).Return(tc.mock.GetOuput.Car , tc.mock.GetOuput.Err).Once()

			gin.SetMode(gin.TestMode)
			router := gin.Default()
			router.Handle(tc.httpMethod, tc.inputPath + tc.inputParamName, handlerInit.getCarByIDGin)
			req, err := http.NewRequest(tc.httpMethod, tc.inputPath + tc.inputParamValue, nil)
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


func TestPostCarsGin(t *testing.T) {
	service := car.ServiceMock{}
	handlerInit := Handler{service: &service}

	type mocks struct {
		AddInput		model.Car
		AddOutput		cars.AddOutput
	}
	tt := []struct {
		name        		string
		mock 				mocks
		httpMethod			string
		inputPath			string
		inputBody			string
		expectedOutput 		string
		expectedHttpCode	int
	}{
		{
			name: "Get_POST_Car_GIN_Success",
			mock: mocks{AddInput: model.Car{Title: "M2", Brand: "BMW", Year: "2020"},
						AddOutput: cars.AddOutput{Car: model.Car{Key: "11c6184d-c848-4848-a7d8-a12e408a4e11" ,Title: "M2", Brand: "BMW", Year: "2020"},
							Err: nil}},
			httpMethod: "POST",
			inputPath: "/carsGin",
			inputBody: "{\n    \"title\": \"M2\",\n    \"brand\": \"BMW\",\n    \"year\": \"2020\"\n}",
			expectedOutput: "{\n    \"key\": \"11c6184d-c848-4848-a7d8-a12e408a4e11\",\n    \"title\": \"M2\",\n    \"brand\": \"BMW\",\n    \"year\": \"2020\"\n}"  ,
			expectedHttpCode:	http.StatusCreated,
		},
		{
			name: "Get_POST_Car_GIN_Invalid_Body",
			mock: mocks{AddInput: model.Car{Title: "M2", Brand: "BMW", Year: "2020"},
				AddOutput: cars.AddOutput{Car: model.Car{Key: "11c6184d-c848-4848-a7d8-a12e408a4e11" ,Title: "M2", Brand: "BMW", Year: "2020"},
					Err: errors.New("this body-request format isn't recognized like a car")}},
			httpMethod: "POST",
			inputPath: "/carsGin",
			inputBody: "{\n    \"title\": \"M2\",\n    \"brand\": \"BMW\",\n    \"year\": \"2020\"",
			expectedOutput: "{}",
			expectedHttpCode:	http.StatusBadRequest,
		},
		{
			name: "Get_POST_Car_GIN_Body_null_values_TITLE",
			mock: mocks{AddInput: model.Car{Title: "", Brand: "BMW", Year: "2020"},
				AddOutput: cars.AddOutput{Car: model.Car{},
					Err: errors.New("obrigatório informar o MODELO do veículo") }},
			httpMethod: "POST",
			inputPath: "/carsGin",
			inputBody: "{\n    \"title\": \"\",\n    \"brand\": \"BMW\",\n    \"year\": \"2020\"\n}",
			expectedOutput: "\"obrigatório informar o MODELO do veículo\"",
			expectedHttpCode:	http.StatusUnprocessableEntity,
		},

	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			service.On("CreateCar", tc.mock.AddInput).Return(tc.mock.AddOutput.Car, tc.mock.AddOutput.Err).Once()
			gin.SetMode(gin.TestMode)
			router := gin.Default()
			router.Handle(tc.httpMethod, tc.inputPath, handlerInit.postCarsGin)

			req, err := http.NewRequest(tc.httpMethod, tc.inputPath, strings.NewReader(tc.inputBody))
			if err != nil {
				t.Fatalf("Couldn't create request: %v\n", err)
			}

			respWriter := httptest.NewRecorder()
			router.ServeHTTP(respWriter, req)

			assert.Equal(t, tc.expectedHttpCode, respWriter.Code)
			assert.Equal(t, tc.expectedOutput, respWriter.Body.String())

		})
	}
}


func TestUpdateCarGin(t *testing.T) {
	service := car.ServiceMock{}
	handlerInit := Handler{service: &service}

	type mocks struct {
		UpdateInput		cars.UpdateInput
		UpdateOutput	cars.UpdateOutput
	}
	tt := []struct {
		name        		string
		mock 				mocks
		httpMethod			string
		inputPath			string
		inputParamName      string
		inputParamValue		string
		inputBody			string
		expectedOutput 		string
		expectedHttpCode	int
	}{
		{
			name: "Update_PATCH_Car_GIN_Success",
			mock: mocks{UpdateInput:
							cars.UpdateInput{Key: "11c6184d-c848-4848-a7d8-a12e408a4e11",
							Car: model.Car{Key: "" ,Title: "M2", Brand: "BMW", Year: "2020"} },
						UpdateOutput: cars.UpdateOutput{
							Car: model.Car{Key: "11c6184d-c848-4848-a7d8-a12e408a4e11" ,Title: "M2", Brand: "BMW", Year: "2020"},
							Err: nil },
			},
			httpMethod: "PATCH",
			inputPath: "/carsGin/",
			inputParamName: ":key",
			inputParamValue: "11c6184d-c848-4848-a7d8-a12e408a4e11",
			inputBody: "{\n    \"title\": \"M2\",\n    \"brand\": \"BMW\",\n    \"year\": \"2020\"\n}",
			expectedOutput: "{\n    \"key\": \"11c6184d-c848-4848-a7d8-a12e408a4e11\",\n    \"title\": \"M2\",\n    \"brand\": \"BMW\",\n    \"year\": \"2020\"\n}"  ,
			expectedHttpCode:	http.StatusOK,
		},

		{
			name: "Update_PATCH_Car_GIN_Error_Another_Brand",
			mock: mocks{UpdateInput:
			cars.UpdateInput{Key: "11c6184d-c848-4848-a7d8-a12e408a4e11",
				Car: model.Car{Key: "" ,Title: "M2", Brand: "BMW", Year: "2020"} },
				UpdateOutput: cars.UpdateOutput{
					Car: model.Car{Key: "11c6184d-c848-4848-a7d8-a12e408a4e11" ,Title: "M2", Brand: "BMW", Year: "2020"},
					Err: errors.New("não é permitido alterar marca do carro. Marca anterior: Mercedez | Marca nova: BMW"),
				},
			},
			httpMethod: "PATCH",
			inputPath: "/carsGin/",
			inputParamName: ":key",
			inputParamValue: "11c6184d-c848-4848-a7d8-a12e408a4e11",
			inputBody: "{\n    \"title\": \"M2\",\n    \"brand\": \"BMW\",\n    \"year\": \"2020\"\n}",
			expectedOutput: "não é permitido alterar marca do carro. Marca anterior: Mercedez | Marca nova: BMW"  ,
			expectedHttpCode:	http.StatusUnprocessableEntity,
		},
		{
			name: "Update_PATCH_Car_GIN_Error_invalid_JSON",
			mock: mocks{UpdateInput:
			cars.UpdateInput{Key: "11c6184d-c848-4848-a7d8-a12e408a4e11",
				Car: model.Car{Key: "" ,Title: "M2", Brand: "BMW", Year: "2020"} },
				UpdateOutput: cars.UpdateOutput{
					Car: model.Car{Key: "11c6184d-c848-4848-a7d8-a12e408a4e11" ,Title: "M2", Brand: "BMW", Year: "2020"},
					Err: nil,
				},
			},
			httpMethod: "PATCH",
			inputPath: "/carsGin/",
			inputParamName: ":key",
			inputParamValue: "11c6184d-c848-4848-a7d8-a12e408a4e11",
			inputBody: "{\n    \"title\": \"M2\",\n    \"brand\": \"BMW\",\n    \"year\":",
			expectedOutput: "this body-request format isn't recognized like a car"  ,
			expectedHttpCode:	http.StatusBadRequest,
		},

	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			service.On("Update", tc.mock.UpdateInput.Key, tc.mock.UpdateInput.Car).Return(tc.mock.UpdateOutput.Car, tc.mock.UpdateOutput.Err).Once()
			gin.SetMode(gin.TestMode)
			router := gin.Default()
			router.Handle(tc.httpMethod, tc.inputPath + tc.inputParamName, handlerInit.updateCarGin)

			req, err := http.NewRequest(tc.httpMethod, tc.inputPath + tc.inputParamValue, strings.NewReader(tc.inputBody))
			if err != nil {
				t.Fatalf("Couldn't create request: %v\n", err)
			}

			respWriter := httptest.NewRecorder()
			router.ServeHTTP(respWriter, req)

			assert.Equal(t, tc.expectedHttpCode, respWriter.Code)
			assert.Equal(t, tc.expectedOutput, respWriter.Body.String())
		})
	}
}




// TEST sem TABLE TEST
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