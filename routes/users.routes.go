package routes

import (
	"encoding/json"
	"net/http"

	"github.com/AgusPane17/go-postgres.git/db"
	"github.com/AgusPane17/go-postgres.git/models"
	"github.com/gorilla/mux"
)

func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	var users []models.User
	db.DB.Find(&users) //No se especifica el tipo de dato que se va a buscar porque lo interpreta en base al tipo de dato del arreglo que es User
	json.NewEncoder(w).Encode(&users)

}

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	
	var user models.User // variable para contener el dato buscado
	params := mux.Vars(r)	//dato que viene por parametro

	db.DB.First(&user, params["id"])
	if user.ID == 0 { // go cuando no encuentra un objeto lo va a rellenar con valores por defecto
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("User not found"))
	}

	db.DB.Model(&user).Association("Tasks").Find(&user.Tasks)

	json.NewEncoder(w).Encode(&user)

}

func PostUserHandler(w http.ResponseWriter, r *http.Request) {
	
	var user models.User //Creo una variable del tipo user
	json.NewDecoder(r.Body).Decode(&user) //Traigo tdo lo que es el body y lo asigno con el puntero de user
	createdUser := db.DB.Create(&user) // Lo guardo en la DB y lo almaceno en una variable
	
	err := createdUser.Error // Manejo errores
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	
	}
	json.NewEncoder(w).Encode(&user) //Devuelvo el objeto creado
}

func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User // variable para contener el dato buscado
	params := mux.Vars(r)	//dato que viene por parametro

	db.DB.First(&user, params["id"])
	if user.ID == 0 { // go cuando no encuentra un objeto lo va a rellenar con valores por defecto
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("User not found"))
	}
	db.DB.Unscoped().Delete(&user) // va a rellenar el campo delete_at de la base de datos, realmente con este comando no lo elimina
	
	// db.DB.Delete(&user) Para eliminarlo realmente se tendria que ocupar este comando
	w.WriteHeader(http.StatusOK)
}