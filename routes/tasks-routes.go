package routes

import (
	"tasks-go/controller"

	"github.com/gorilla/mux"
)

var RegisterTasksRoutes = func(router *mux.Router) {
	router.HandleFunc("/v1/tasks", controller.CreateTask).Methods("POST")
	router.HandleFunc("/v1/tasks", controller.GetAllTasks).Methods("GET")
	router.HandleFunc("/v1/tasks/{id}", controller.GetTaskById).Methods("GET")
	router.HandleFunc("/v1/tasks/{id}", controller.UpdateTask).Methods("PUT")
	router.HandleFunc("/v1/tasks/{id}", controller.DeleteTask).Methods("DELETE")
}
