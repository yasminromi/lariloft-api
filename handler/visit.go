package handler

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/yasminromi/lariloft-api/model"
)

func (as *ApiServer) GetVisitHandle(w http.ResponseWriter, r *http.Request) {

	enableCors(&w)

	vars := mux.Vars(r)
	id := vars["id"]
	idInt, _ := strconv.Atoi(id)

	visit, err := as.DB.GetVisit(idInt)
	if err != nil {
		log.Println("error on get user", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(visit); err != nil {
		log.Println("error encode user object", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
}

func (as *ApiServer) GetAllVisitsHandle(w http.ResponseWriter, r *http.Request) {

	enableCors(&w)

	vars := mux.Vars(r)
	userID := vars["userId"]

	visits, err := as.DB.GetAllVisitApartmentsByUser(userID)
	if err != nil {
		log.Println("error on get visits", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(visits); err != nil {
		log.Println("error encode user object", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
}

func (as *ApiServer) CreateVisitHandle(w http.ResponseWriter, r *http.Request) {

	enableCors(&w)

	//Leitura do body da requisição
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		log.Println("Error on getting body content", err)
		http.Error(w, "Error on getting body content", 500)
		return
	}
	visit := &model.Visit{}
	err = json.Unmarshal(b, &visit)
	if err != nil {
		log.Println("Error on unmarshal info from body", err)
		http.Error(w, "Error on unmarshal info from body", 500)
		return
	}

	if err := as.DB.CreateVisit(visit); err != nil {
		log.Println("error on create user data", err)
		http.Error(w, "Error on unmarshal info from body", 500)
		return
	}
}
