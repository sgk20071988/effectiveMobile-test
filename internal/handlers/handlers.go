package handlers

import (
	context "effectiveMobileTest/internal/context"
	"effectiveMobileTest/internal/db"
	"effectiveMobileTest/internal/model"
	repository "effectiveMobileTest/internal/repository"
	"encoding/json"
	"net/http"
	"strconv"
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
	var limit, offset int
	limit, err = strconv.Atoi(params["limit"])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	offset, err = strconv.Atoi(params["offset"])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	var cars []model.Car
	cars, err = carRep.GetCars(limit, offset)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(cars)
	w.WriteHeader(http.StatusOK)
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
