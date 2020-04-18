package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

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

	w.Header().Set("Location", fmt.Sprintf("%s%s/%s", r.Host, r.RequestURI, user.ID))
	responses.JSON(w, http.StatusCreated, user)
}

func GetUser(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	uid, err := strconv.ParseInt(params["id"], 10, 32)

	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	user, err := services.User.FindById(uid)
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
	uid, err := strconv.ParseInt(params["id"], 10, 32)

	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	updatedCount, err := services.User.Update(uid, user)

	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	responses.JSON(w, http.StatusAccepted, updatedCount)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	uid, err := strconv.ParseInt(params["id"], 10, 32)

	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	deleteCount, err := services.User.Delete(uid)

	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	w.Header().Set("Entity", fmt.Sprintf("%s", n))
	responses.JSON(w, http.StatusNoContent, deleteCount)
}
