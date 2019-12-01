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

func (as *ApiServer) GetAgentHandle(w http.ResponseWriter, r *http.Request) {

	enableCors(&w)

	vars := mux.Vars(r)
	id := vars["id"]
	idInt, _ := strconv.Atoi(id)

	agent, err := as.DB.GetAgent(idInt)
	if err != nil {
		log.Println("error on get user", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(agent); err != nil {
		log.Println("error encode user object", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
}

func (as *ApiServer) GetAllAgentsHandle(w http.ResponseWriter, r *http.Request) {

	enableCors(&w)

	agents, err := as.DB.GetAllAgents()
	if err != nil {
		log.Println("error on get agents", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(agents); err != nil {
		log.Println("error encode user object", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
}

func (as *ApiServer) CreateAgentHandle(w http.ResponseWriter, r *http.Request) {

	enableCors(&w)

	//Leitura do body da requisição
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		log.Println("Error on getting body content", err)
		http.Error(w, "Error on getting body content", 500)
		return
	}
	agent := &model.Agent{}
	err = json.Unmarshal(b, &agent)
	if err != nil {
		log.Println("Error on unmarshal info from body", err)
		http.Error(w, "Error on unmarshal info from body", 500)
		return
	}

	if err := as.DB.CreateAgent(agent); err != nil {
		log.Println("error on create user data", err)
		http.Error(w, "Error on unmarshal info from body", 500)
		return
	}
}
