package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/solrac87/rest/src/api/models"
	"github.com/solrac87/rest/src/api/responses"
	"github.com/solrac87/rest/src/api/services"
)

func CreatePost(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	post := models.Post{}
	err = json.Unmarshal(body, &post)

	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
	}

	post, err = services.Post.CreatePost(post)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	w.Header().Set("Location", fmt.Sprintf("%s%s/%s", r.Host, r.RequestURI, post.ID))
	responses.JSON(w, http.StatusCreated, post)
}
