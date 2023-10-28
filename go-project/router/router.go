package router

import (
	"fmt"
	"go-project/controllers"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

const baseApiURL = "/api/v1"

func Routes() http.Handler {
	router := chi.NewRouter()

	// specify who is allowed to connect
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	router.Use(middleware.Heartbeat("/ping"))

	// USER ENDPOINTS
	usersBaseString := fmt.Sprintf("%s/users", baseApiURL)
	router.Get(usersBaseString, controllers.GetAllUsers)
	router.Post(usersBaseString, controllers.CreateUser)
	router.Put(usersBaseString, controllers.UpdateUser)

	userProjectsString := fmt.Sprintf("%s/projects", usersBaseString)
	router.Get(userProjectsString, controllers.GetAllUserProjects)

	singleUserBaseString := fmt.Sprintf("%s/{email}", usersBaseString)
	router.Post(singleUserBaseString, controllers.LoginUser)

	// PROJECT ENDPOINTS
	projectsBaseString := fmt.Sprintf("%s/projects", baseApiURL)
	router.Get(projectsBaseString, controllers.GetAllProjects)
	router.Post(projectsBaseString, controllers.CreateProject)
	router.Put(projectsBaseString, controllers.UpdateProject)

	projectTasksString := fmt.Sprintf("%s/{id}/tasks", projectsBaseString)
	router.Get(projectTasksString, controllers.GetAllProjectTasks)
	router.Post(projectTasksString, controllers.CreateTask)

	// TASKS ENDPOINTS
	// tasksBaseString := fmt.Sprintf("%s/tasks", baseApiURL)
	// router.Get(tasksBaseString, controllers.GetAllTasks)
	// router.Put(tasksBaseString, controllers.UpdateTask)
	// router.Delete(tasksBaseString, controllers.DeleteTask)

	return router
}
