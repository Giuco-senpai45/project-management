package controllers

import (
	"go-project/helpers"
	"go-project/models"
	"net/http"
)

func GetAllSubTasks(w http.ResponseWriter, r *http.Request) {
	subtasks, err := subTaskModel.GetAllSubTasks()
	if err != nil {
		helpers.ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}
	helpers.WriteJSON(w, http.StatusOK, subtasks)
}

func CreateSubTask(w http.ResponseWriter, r *http.Request) {
	var subtaskResp models.SubTask

	err := helpers.ReadJSON(w, r, &subtaskResp)
	if err != nil {
		helpers.ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}

	createdSubTask, err := subTaskModel.CreateSubTask(&subtaskResp)
	if err != nil {
		helpers.ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}
	helpers.WriteJSON(w, http.StatusOK, createdSubTask)
}

func UpdateSubTask(w http.ResponseWriter, r *http.Request) {
	var subtaskResp models.SubTask

	err := helpers.ReadJSON(w, r, &subtaskResp)
	if err != nil {
		helpers.ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}

	updatedSubTask, err := subTaskModel.UpdateSubTask(&subtaskResp)
	if err != nil {
		helpers.ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}
	helpers.WriteJSON(w, http.StatusOK, updatedSubTask)
}
