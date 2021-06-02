package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"tasks-go/models"
	"tasks-go/utils"

	"github.com/gorilla/mux"
)

func CreateTask(w http.ResponseWriter, r *http.Request) {
	CreateTask := &models.Task{}
	utils.ParseBody(r, CreateTask)
	t := CreateTask.CreateTask()
	var Resp struct {
		Id int64 `json:"id"`
	}
	Resp.Id = t.Id
	res, _ := json.Marshal(&Resp)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(res)
}

func GetAllTasks(w http.ResponseWriter, r *http.Request) {
	tasks := models.GetAllTasks()
	var Resp struct {
		Tasks []models.Task `json:"tasks"`
	}
	Resp.Tasks = tasks
	res, _ := json.Marshal(&Resp)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetTaskById(w http.ResponseWriter, r *http.Request) {
	Id := mux.Vars(r)["id"]
	ID, err := strconv.ParseInt(Id, 0, 0)
	if err != nil {
		fmt.Println("Error while Parsing")
	}
	taskDetails, _ := models.GetTaskById(ID)
	w.Header().Set("Content-Type", "application/json")
	var ErrorResp struct {
		Error string `json:"error"`
	}
	if taskDetails.Id < 1 {
		ErrorResp.Error = "There is no task at that Id"
		resp, _ := json.Marshal(ErrorResp)
		w.WriteHeader(http.StatusNotFound)
		w.Write(resp)
		return
	}
	res, _ := json.Marshal(taskDetails)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	var updateTask = &models.Task{}
	utils.ParseBody(r, updateTask)
	Id := mux.Vars(r)["id"]
	ID, err := strconv.ParseInt(Id, 0, 0)
	if err != nil {
		fmt.Println("Error while Parsing")
	}
	taskDetails, db := models.GetTaskById(ID)
	w.Header().Set("Content-Type", "application/json")
	var ErrorResp struct {
		Error string `json:"error"`
	}
	if taskDetails.Id < 1 {
		ErrorResp.Error = "There is no task at that Id"
		resp, _ := json.Marshal(ErrorResp)
		w.WriteHeader(http.StatusNotFound)
		w.Write(resp)
		return
	}

	if updateTask.Title != "" {
		taskDetails.Title = updateTask.Title
	}
	taskDetails.IsCompleted = updateTask.IsCompleted
	if updateTask.Notify != "" {
		taskDetails.Notify = updateTask.Notify
	}
	db.Save(&taskDetails)
	res, _ := json.Marshal(taskDetails)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	Id := mux.Vars(r)["id"]
	ID, err := strconv.ParseInt(Id, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	models.DeleteTask(ID)
	w.WriteHeader(http.StatusNoContent)
}
