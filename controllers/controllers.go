package controllers

import (
	"encoding/json"
	"net/http"
	"reactgoboilerplate/models"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	SetHeaders(w)
	w.Header().Set("Access-Control-Allow-Methods", "GET")

	userOne := models.User{
		FirstName: "John",
		LastName: "Doe",
		Email: "jdoe@gmail.com",
	}

	userTwo := models.User{
		FirstName: "Jane",
		LastName: "Smith",
		Email: "jsmith@gmail.com",
	}

	results := []models.User{userOne,userTwo}
	json.NewEncoder(w).Encode(results)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	SetHeaders(w)
	w.Header().Set("Access-Control-Allow-Methods", "POST")

	var user models.User
	_ = json.NewDecoder(r.Body).Decode(&user)

	if user.Email == "" {
		http.Error(w, "Email Is Required", http.StatusBadRequest)
		return
	}
	if user.FirstName == "" {
		http.Error(w, "First Name Is Required", http.StatusBadRequest)
		return
	}
	if user.LastName == "" {
		http.Error(w, "Last Name Is Required", http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(user)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	SetHeaders(w)
	w.Header().Set("Access-Control-Allow-Methods", "GET")

	user := models.User{
		FirstName: "John",
		LastName: "Doe",
		Email: "jdoe@gmail.com",
	}

	json.NewEncoder(w).Encode(user)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	SetHeaders(w)
	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	// Get Body
	var user models.User
	_ = json.NewDecoder(r.Body).Decode(&user)

	// Simulate update by returning body
	json.NewEncoder(w).Encode(user)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	SetHeaders(w)
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")
	// Get Body
	type DeletedBody struct {
		Message string `json:"message"`
	}

	// Simulate delete by returning deleted message
	newDeleteBody := DeletedBody{
		Message: "User Deleted",
	}

	json.NewEncoder(w).Encode(newDeleteBody)

}