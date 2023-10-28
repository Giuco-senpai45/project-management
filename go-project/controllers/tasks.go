package controllers

import (
	"go-project/helpers"
	"go-project/middleware"
	"go-project/models"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func GetAllTasks(w http.ResponseWriter, r *http.Request) {
	tasks, err := taskModel.GetAllTasks()
	if err != nil {
		helpers.ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}
	helpers.WriteJSON(w, http.StatusOK, tasks)
}

func CreateTask(w http.ResponseWriter, r *http.Request) {
	var taskResp models.Task

	_, err := middleware.RequireAuth(w, r)
	if err != nil {
		helpers.ErrorJSON(w, err, http.StatusUnauthorized)
		return
	}

	projectId, _ := strconv.Atoi(chi.URLParam(r, "id"))
	project, err := projectModel.GetProjectByID(int64(projectId))
	if err != nil {
		helpers.ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}

	err = helpers.ReadJSON(w, r, &taskResp)
	if err != nil {
		helpers.ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}

	createdTask, err := taskModel.CreateTask(&taskResp, project)
	if err != nil {
		helpers.ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}
	helpers.WriteJSON(w, http.StatusOK, createdTask)
}

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	var taskResp models.Task

	err := helpers.ReadJSON(w, r, &taskResp)
	if err != nil {
		helpers.ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}

	updatedTask, err := taskModel.UpdateTask(&taskResp)
	if err != nil {
		helpers.ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}
	helpers.WriteJSON(w, http.StatusOK, updatedTask)
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	var taskResp models.Task

	err := helpers.ReadJSON(w, r, &taskResp)
	if err != nil {
		helpers.ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}

	err = taskModel.DeleteTask(&taskResp)
	if err != nil {
		helpers.ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}
	helpers.WriteJSON(w, http.StatusOK, taskResp)
}
