package handlers

import (
	context "effectiveMobileTest/internal/context"
	"effectiveMobileTest/internal/model"
	repository "effectiveMobileTest/internal/repository"
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func GetCars(w http.ResponseWriter, r *http.Request) {
	// swagger:route GET /cars cars listCar
	//
	// Lists of cars with pagination.
	//
	// This will show all recorded people.
	//
	//     Consumes:
	//     - application/json
	//
	//     Produces:
	//     - application/json
	//
	//     Schemes: http, https
	//
	//     Params:
	//     limit: limit
	//     offset: offset
	//     filters: filters
	//
	//     Responses:
	//       200: carsResponse
	//       400: bad request
	//       500: bad connection
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	filters := map[string]string{}
	err := json.Unmarshal([]byte(r.URL.Query().Get("filters")), &filters)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	params := map[string]string{}
	limit := r.URL.Query().Get("limit")
	if len(limit) > 0 {
		params["limit"] = limit
	}
	offset := r.URL.Query().Get("offset")
	if len(offset) > 0 {
		params["offset"] = offset
	}
	db, err := context.GetDB()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	carRep := repository.Repository{DB: db}
	var cars []model.Car
	cars, err = carRep.GetCars(params, filters)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(cars)
	w.WriteHeader(http.StatusOK)
}

func InsertCar(w http.ResponseWriter, r *http.Request) {
	// swagger:route POST /cars/ cars listCars
	//
	// Insert new car in cars list.
	//
	//     Consumes:
	//     - application/json
	//
	//     Produces:
	//     - application/json
	//
	//     Schemes: http, https
	//
	//     Params:
	//
	//     Responses:
	//       200: carResponse
	//       400: bad request
	//       500: bad connection
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	decoder := json.NewDecoder(r.Body)
	var car model.Car
	err := decoder.Decode(&car)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	db, err := context.GetDB()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	carRep := repository.Repository{DB: db}
	err = carRep.Insert(car)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func UpdateCar(w http.ResponseWriter, r *http.Request) {
	// swagger:route PATCH /cars/{regNum} cars listCars
	//
	// Update row with identified registration number.
	//
	//     Consumes:
	//     - application/json
	//
	//     Produces:
	//     - application/json
	//
	//     Schemes: http, https
	//
	//     Params:
	//       regNum: regNum
	//
	//     Responses:
	//       200: carResponse
	//       400: bad request
	//       500: bad connection
	params := mux.Vars(r)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	regNum, ok := params["regNum"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	data := map[string]string{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&data)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	db, err := context.GetDB()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	carRep := repository.Repository{DB: db}
	err = carRep.Update(regNum, params)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func DeleteCar(w http.ResponseWriter, r *http.Request) {
	// swagger:route DELETE /cars/{regNum} cars listCars
	//
	// Delete row with identified registration number.
	//
	//     Consumes:
	//     - application/json
	//
	//     Produces:
	//     - application/json
	//
	//     Schemes: http, https
	//
	//     Params:
	//       regNum: regNum
	//
	//     Responses:
	//       200: carResponse
	//       400: bad request
	//       500: bad connection
	params := mux.Vars(r)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	regNum, ok := params["regNum"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	db, err := context.GetDB()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	carRep := repository.Repository{DB: db}
	err = carRep.DeleteCar(regNum)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func GetCar(w http.ResponseWriter, r *http.Request) {
	// swagger:route GET /cars/{regNum} cars listCar
	//
	// Lists of cars with pagination.
	//
	// This will show the record of an identified registration number.
	//
	//     Consumes:
	//     - application/json
	//
	//     Produces:
	//     - application/json
	//
	//     Schemes: http, https
	//
	//     Params:
	//       regNum: regNum
	//
	//     Responses:
	//       200: carResponse
	//       404: jsonError
	//       500: bad connection
	params := mux.Vars(r)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	db, err := context.GetDB()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	carRep := repository.Repository{DB: db}
	car, err := carRep.GetCar(params["regNum"])

	if err != nil {
		// If we didn't find it, 404
		w.WriteHeader(http.StatusNotFound)
		if err := json.NewEncoder(w).Encode(jsonError{Message: "Not Found"}); err != nil {
			panic(err)
		}
	}

	w.WriteHeader(http.StatusOK)
	time.Sleep(1000 * time.Millisecond)
	if err := json.NewEncoder(w).Encode(car); err != nil {
		panic(err)
	}
}
