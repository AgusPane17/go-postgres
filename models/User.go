package models

import "gorm.io/gorm"

type User struct {
	gorm.Model

	FirstName string `gorm:"not null" json:"firstName"`
	LastName  string `gorm:"not null" json:"lastName"` 
	Email     string `gorm:"not null;unique_index" json:"email"`
	Tasks     []Task `json:"task"`
}

// json:"...": Etiqueta utilizada para personalizar la 
// serialización/deserialización de la estructura a/desde 
// formato JSON. Establece los nombres de los campos en el objeto JSON.