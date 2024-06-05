package handeler

import (
	"encoding/json"
	"net/http"
	"taqsir/model"
	"taqsir/utils"

	"github.com/gorilla/mux"
)

var users = make(map[string]string)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user model.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid request payload"))
		return
	}

	if _, ok := users[user.ID]; ok {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("User already exists"))
		return
	}

	users[user.ID] = user.Data

	cation := model.Cation{
		UserID: user.ID,
		Action: "create",
	}
	utils.SendNotification(cation)

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("User created"))
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var user model.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid request payload"))
		return
	}

	if _, ok := users[id]; !ok {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("User not found"))
		return
	}

	users[id] = user.Data

	cation := model.Cation{
		UserID: id,
		Action: "update",
	}
	utils.SendNotification(cation)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("User updated"))
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	if _, exists := users[id]; !exists {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("User not found"))
		return
	}

	delete(users, id)

	cation := model.Cation{
		UserID: id,
		Action: "delete",
	}
	utils.SendNotification(cation)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("User deleted"))
}
