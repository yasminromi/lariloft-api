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

func (as *ApiServer) GetApartmentHandle(w http.ResponseWriter, r *http.Request) {

	enableCors(&w)

	vars := mux.Vars(r)
	id := vars["id"]
	idInt, _ := strconv.Atoi(id)

	apartment, err := as.DB.GetApartment(idInt)
	if err != nil {
		log.Println("error on get user", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(apartment); err != nil {
		log.Println("error encode user object", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
}

func (as *ApiServer) GetAllApartmentsHandle(w http.ResponseWriter, r *http.Request) {

	enableCors(&w)

	apartments, err := as.DB.GetAllApartments()
	if err != nil {
		log.Println("error on get users", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(apartments); err != nil {
		log.Println("error encode user object", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
}

func (as *ApiServer) CreateApartmentHandle(w http.ResponseWriter, r *http.Request) {

	enableCors(&w)

	//Leitura do body da requisição
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		log.Println("Error on getting body content", err)
		http.Error(w, "Error on getting body content", 500)
		return
	}
	apartment := &model.Apartment{}
	err = json.Unmarshal(b, &apartment)
	if err != nil {
		log.Println("Error on unmarshal info from body", err)
		http.Error(w, "Error on unmarshal info from body", 500)
		return
	}

	if err := as.DB.CreateApartment(apartment); err != nil {
		log.Println("error on create user data", err)
		http.Error(w, "Error on unmarshal info from body", 500)
		return
	}
}
