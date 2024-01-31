package main

import (
	"log"
	"net/http"

	"github.com/AgusPane17/go-postgres.git/db"
	"github.com/AgusPane17/go-postgres.git/models"
	"github.com/AgusPane17/go-postgres.git/routes"
	"github.com/gorilla/mux"
)



func main() {

	db.ConectionDB()


	err := db.DB.AutoMigrate(&models.Task{}, &models.User{})// Gorm utiliza las estructuras proporcionadas para determinar las tablas correspondientes en la base de datos y crearlas.
	if err != nil {
    	log.Fatal(err)
	} // En estas lineas se importa la tabla y la estructura para que genere las tablas con el automigrate 

	r := mux.NewRouter()

	r.HandleFunc("/", routes.HomeHandler)
	
	// USERS ROUTES
	r.HandleFunc("/users", routes.GetUsersHandler).Methods("GET")
	r.HandleFunc("/users/{id}", routes.GetUserHandler).Methods("GET")
	r.HandleFunc("/users", routes.PostUserHandler).Methods("POST")
	r.HandleFunc("/users/{id}", routes.DeleteUserHandler).Methods("DELETE")

	// TASKS ROUTES
	r.HandleFunc("/tasks", routes.GetTasksHandler).Methods("GET")
	r.HandleFunc("/tasks/{id}", routes.GetTaskHandler).Methods("GET")
	r.HandleFunc("/tasks", routes.PostTaskHandler).Methods("POST")
	r.HandleFunc("/tasks/{id}", routes.DeleteTaskHandler).Methods("DELETE")

	
	http.ListenAndServe(":8080", r)
}