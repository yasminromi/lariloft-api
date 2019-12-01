package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/yasminromi/lariloft-api/model"
	"github.com/yasminromi/lariloft-api/service"
)

type Handler struct {
	Service *service.ElasticService
}

type ApiServer struct {
	DB model.LariLoftDB
}

func (h *Handler) SendViaPost(w http.ResponseWriter, r *http.Request) {

	enableCors(&w)

	if r.Method == "POST" {

		var payload model.Interest

		if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
			log.Printf("ERROR: %s", err)
			http.Error(w, "Bad request", http.StatusTeapot)
			return
		}

		defer r.Body.Close()

		log.Printf("payload sending via post: %v", payload)

		ctx := r.Context()

		error := h.Service.SaveToElastic(ctx, payload)
		if error != nil {
			log.Printf("error saving to ES: %v", error)
			return
		}

	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}
