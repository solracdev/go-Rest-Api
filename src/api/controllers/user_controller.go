package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/solrac87/rest/src/api/models"
	"github.com/solrac87/rest/src/api/responses"
	"github.com/solrac87/rest/src/api/services"
)

type UserController struct {
}

func GetUsers(w http.ResponseWriter, r *http.Request) {

	users, err := services.User.FindAll()

	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, users)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	user := models.User{}
	err = json.Unmarshal(body, &user)

	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
	}

	user, err = services.User.CreateUser(user)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	w.Header().Set("Location", fmt.Sprintf("%s%s/%s", r.Host, r.RequestURI, user.NickName))
	responses.JSON(w, http.StatusCreated, user)
}

func GetUser(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	nickname := params["nick"]
	if nickname == "" {
		responses.ERROR(w, http.StatusBadRequest, errors.New("Nickname is required"))
	}

	user, err := services.User.FindByNickname(nickname)
	if err != nil {
		responses.ERROR(w, http.StatusNotFound, err)
		return
	}

	responses.JSON(w, http.StatusOK, user)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	user := models.User{}
	err = json.Unmarshal(body, &user)

	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
	}

	params := mux.Vars(r)
	n := params["nick"]

	updatedCount, err := services.User.Update(n, user)

	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	responses.JSON(w, http.StatusAccepted, updatedCount)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	n := params["nick"]

	deleteCount, err := services.User.Delete(n)

	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	w.Header().Set("Entity", fmt.Sprintf("%s", n))
	responses.JSON(w, http.StatusNoContent, deleteCount)
}
