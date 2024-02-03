package controllers

import (
	dbmysql "backend/src/db_mysql"
	"backend/src/models"
	"backend/src/repositories"
	"backend/src/response"
	"encoding/json"
	"io"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	bodyRequest, err := io.ReadAll(r.Body)
	if err != nil {
		response.Err(w, http.StatusUnprocessableEntity, err)
		return
	}

	var user models.User
	if err = json.Unmarshal(bodyRequest, &user); err != nil {
		response.Err(w, http.StatusBadRequest, err)
		return
	}

	if err = user.PrepareString(); err != nil {
		response.Err(w, http.StatusBadRequest, err)
		return
	}

	db, err := dbmysql.ConnectMySql()
	if err != nil {
		response.Err(w, http.StatusInternalServerError, err)
		return
	}

	defer db.Close()

	repository := repositories.NewUsersRepository(db)
	user.Id, err = repository.CreateUser(user)
	if err != nil {
		response.Err(w, http.StatusInternalServerError, err)
		return
	}

	response.ToJson(w, http.StatusCreated, user)

}

func FindAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Finding all users!"))
}

func FindUserById(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Finding user!"))
}

func UpdateUserById(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Updating user!"))
}

func DeleteUserById(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Removing user!"))
}
