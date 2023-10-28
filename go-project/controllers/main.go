package controllers

import "go-project/models"

var modelsWrapper models.Models
var userModel = modelsWrapper.User
var projectModel = modelsWrapper.Project
var taskModel = modelsWrapper.Task
var subTaskModel = modelsWrapper.SubTask
