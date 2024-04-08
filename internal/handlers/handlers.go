package handlers

import (
	context "effectiveMobileTest/internal/context"
	"effectiveMobileTest/internal/db"
	"effectiveMobileTest/internal/model"
	repository "effectiveMobileTest/internal/repository"
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func GetCars(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	params := mux.Vars(r)
	db, err := context.GetDB()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	carRep := repository.Repository{DB: db}
	var cars []model.Car
	cars, err = carRep.GetCars(params)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(cars)
	w.WriteHeader(http.StatusOK)
}

func InsertCar(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	params := mux.Vars(r)
	db, err := context.GetDB()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	carRep := repository.Repository{DB: db}
	car := model.Car{}
	regNum, ok := params["regNum"]
	if ok {
		car.RegNum = regNum
	}
	owner, ok := params["owner"]
	if ok {
		car.Owner = owner
	}
	model, ok := params["model"]
	if ok {
		car.Model = model
	}
	mark, ok := params["mark"]
	if ok {
		car.Mark = mark
	}
	err = carRep.Insert(car)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func UpdateCar(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	params := mux.Vars(r)
	db, err := context.GetDB()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	carRep := repository.Repository{DB: db}
	carRep.Update(params)
}

func GetCar(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	for _, item := range db.Get() {
		if item.RegNum == params["regNum"] {
			w.WriteHeader(http.StatusOK)
			// add a arbitraty pause of 1 second
			time.Sleep(1000 * time.Millisecond)
			if err := json.NewEncoder(w).Encode(item); err != nil {
				panic(err)
			}
			return
		}
	}
	// If we didn't find it, 404
	w.WriteHeader(http.StatusNotFound)
	if err := json.NewEncoder(w).Encode(jsonError{Message: "Not Found"}); err != nil {
		panic(err)
	}
}
