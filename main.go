package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/yasminromi/lariloft-api/handler"
	"github.com/yasminromi/lariloft-api/model"
	"github.com/yasminromi/lariloft-api/service"
	"gopkg.in/olivere/elastic.v6"
)

func main() {

	routes := mux.NewRouter()

	db := model.WeeHackDB{}
	//Inicialização do repositorio
	db.MustInit()

	if err := godotenv.Load(); err != nil {
		log.Println("File .env not found, reading configuration from ENV")
	}

	elasticClient, err := elastic.NewClient(elastic.SetURL(model.ElasticSearch()), elastic.SetSniff(false))
	if err != nil {
		log.Fatal("Error Creating Elastic Client: ", err)
	}

	log.Printf("Elastic Search Client Created")

	elasticService := &service.ElasticService{
		ElasticCLI: elasticClient,
	}

	chatHandler := &handler.Handler{
		Service: elasticService,
	}

	apiServer := handler.ApiServer{
		DB: db,
	}

	//Rotas de consulta
	routes.HandleFunc("/api/visit/{id:[0-9]+}", apiServer.GetVisitHandle).Methods("GET")
	routes.HandleFunc("/api/visit/all", apiServer.GetAllVisitsHandle).Methods("GET")

	routes.HandleFunc("/api/apartment/{id:[0-9]+}", apiServer.GetApartmentHandle).Methods("GET")
	routes.HandleFunc("/api/apartment/{userId:[0-9]+}", apiServer.GetVisitApartmentsByUserHandle).Methods("GET")

	routes.HandleFunc("/api/agent/{id:[0-9]+}", apiServer.GetAgentHandle).Methods("GET")
	routes.HandleFunc("/api/agent/all", apiServer.GetAllAgentsHandle).Methods("GET")

	routes.HandleFunc("/api/user/{id:[0-9]+}", apiServer.GetUserHandle).Methods("GET")
	routes.HandleFunc("/api/user/all", apiServer.GetAllUsersHandle).Methods("GET")

	//Rotas de criação
	routes.HandleFunc("/api/visit", apiServer.CreateVisiyHandle).Methods("POST")

	routes.HandleFunc("/api/login/{name}", apiServer.LoginUserHandle).Methods("POST")

	routes.HandleFunc("/api/agent", apiServer.CreateAgentHandle).Methods("POST")

	routes.HandleFunc("/api/apartment", apiServer.CreateApartmentHandle).Methods("POST")

	routes.HandleFunc("/api/user", apiServer.CreateUserHandle).Methods("POST")

	http.Handle("/", routes)
	http.HandleFunc("/sendIntention", chatHandler.SendViaPost)

	go chatHandler.HandleMessages()

	log.Println("http server started on " + os.Getenv("PORT"))
	log.Println("database started on " + model.Connection())
	log.Println("bonsai started on " + model.ElasticSearch())

	error := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	if error != nil {
		log.Fatal("ListenAndServe: ", error)
	}

}
