package controllers

import (
	dbmysql "backend/src/db_mysql"
	"backend/src/models"
	"backend/src/repositories"
	"backend/src/response"
	"encoding/json"
	"io"
	"net/http"
	"strings"
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
	nameOrNick := strings.ToLower(r.URL.Query().Get("user"))
	db, erro := dbmysql.ConnectMySql()
	if erro != nil {
		response.Err(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repository := repositories.NewUsersRepository(db)
	user, erro := repository.FindByNameOrNick(nameOrNick)
	if erro != nil {
		response.Err(w, http.StatusInternalServerError, erro)
		return
	}

	response.ToJson(w, http.StatusOK, user)
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
