package controllers

import (
	"go-project/helpers"
	"go-project/middleware"
	"go-project/models"
	"net/http"
)

func GetAllProjects(w http.ResponseWriter, r *http.Request) {
	all, err := projectModel.GetAllProjects()
	if err != nil {
		helpers.ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}
	helpers.WriteJSON(w, http.StatusOK, all)
}

func CreateProject(w http.ResponseWriter, r *http.Request) {
	var projectResp models.Project

	user, err := middleware.RequireAuth(w, r)
	if err != nil {
		helpers.ErrorJSON(w, err, http.StatusUnauthorized)
		return
	}

	err = helpers.ReadJSON(w, r, &projectResp)
	if err != nil {
		helpers.ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}

	createdProject, err := projectModel.CreateProject(&projectResp, user)
	if err != nil {
		helpers.ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}
	helpers.WriteJSON(w, http.StatusOK, createdProject)
}

func UpdateProject(w http.ResponseWriter, r *http.Request) {
	var projectResp models.Project

	_, err := middleware.RequireAuth(w, r)
	if err != nil {
		helpers.ErrorJSON(w, err, http.StatusUnauthorized)
		return
	}

	err = helpers.ReadJSON(w, r, &projectResp)
	if err != nil {
		helpers.ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}

	updatedProject, err := projectModel.UpdateProject(&projectResp)
	if err != nil {
		helpers.ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}
	helpers.WriteJSON(w, http.StatusOK, updatedProject)
}

func GetAllProjectTasks(w http.ResponseWriter, r *http.Request) {
	var projectResp models.Project

	_, err := middleware.RequireAuth(w, r)
	if err != nil {
		helpers.ErrorJSON(w, err, http.StatusUnauthorized)
		return
	}

	err = helpers.ReadJSON(w, r, &projectResp)
	if err != nil {
		helpers.ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}

	project, err := projectModel.GetAllProjectTasks(&projectResp)
	if err != nil {
		helpers.ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}
	helpers.WriteJSON(w, http.StatusOK, project)
}
