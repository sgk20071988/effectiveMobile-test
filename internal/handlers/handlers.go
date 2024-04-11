package handlers

import (
	_ "effectiveMobileTest/cmd/api-server/docs"
	context "effectiveMobileTest/internal/context"
	"effectiveMobileTest/internal/model"
	repository "effectiveMobileTest/internal/repository"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func GetCars(w http.ResponseWriter, r *http.Request) {
	fmt.Println("get cars")
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
	fmt.Println("insert car")
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
	fmt.Println("update car")
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
	fmt.Println("delete car")
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

// GetCar godoc
// @Summary      Show car
// @Description  get string by regNum
// @Tags         cars
// @Accept       json
// @Produce      json
// @Param        regNum   path      int  true  "Car registration number"
// @Success      200  {object}  model.Car
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /cars/{regNum} [get]
func GetCar(w http.ResponseWriter, r *http.Request) {
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
		w.WriteHeader(http.StatusNotFound)
		if err := json.NewEncoder(w).Encode(jsonError{Message: "Not Found"}); err != nil {
			panic(err)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(car); err != nil {
		panic(err)
	}
}
