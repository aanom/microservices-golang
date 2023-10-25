package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/backend-services/controllers/dataservice"
	response "github.com/backend-services/models/response"
	"github.com/backend-services/utils/constants"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

func CreateBlog(w http.ResponseWriter, r *http.Request) {
	var result = response.HealthCheckResponse{}
	result.Health = true
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	data, _ := json.Marshal(result)
	w.Write(data)
	return
}

func GetBlogByID(w http.ResponseWriter, r *http.Request) {
	logHeader := "[Enterprise] Get User API "
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	ID := params["id"]

	if ID == "" {
		RespondWithError(w, http.StatusBadRequest, constants.MessagesInvalidBlockId, constants.InvalidParams)
		return
	}

	blogDB, dbRes := dataservice.DataService.GetBlogByID(ID)
	logrus.Error(logHeader, blogDB, dbRes)
	if dbRes != nil {
		RespondWithError(w, http.StatusUnauthorized, "blog not exist", constants.FetchDBDataError)
		return
	}

	var blog response.Blog
	blog.ID = blogDB.ID.String
	blog.Title = blogDB.Title.String
	blog.Description = blogDB.Description.String
	blog.CreatedAt = blogDB.CreatedAt.Time
	blog.UpdatedAt = blogDB.UpdatedAt.Time

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(blog)
	return
}

func UpdateBlog(w http.ResponseWriter, r *http.Request) {
	var result = response.HealthCheckResponse{}
	result.Health = true
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	data, _ := json.Marshal(result)
	w.Write(data)
	return
}

func DeleteBlogByID(w http.ResponseWriter, r *http.Request) {
	var result = response.HealthCheckResponse{}
	result.Health = true
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	data, _ := json.Marshal(result)
	w.Write(data)
	return
}

func GetBlogs(w http.ResponseWriter, r *http.Request) {
	var result = response.HealthCheckResponse{}
	result.Health = true
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	data, _ := json.Marshal(result)
	w.Write(data)
	return
}
