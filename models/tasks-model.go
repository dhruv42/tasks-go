package models

import (
	"tasks-go/config"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Task struct {
	Id          int64  `json:"id"`
	Title       string `json:"title"`
	IsCompleted bool   `json:"is_completed"`
	Notify      string `json:"notify"`
}

func init() {
	config.DbOpen()
	db = config.GetDB()
	db.AutoMigrate(&Task{})
}

func (t *Task) CreateTask() *Task {
	db.NewRecord(t)
	db.Create(&t)
	return t
}

func GetAllTasks() []Task {
	var Tasks []Task
	db.Find(&Tasks)
	return Tasks
}

func GetTaskById(Id int64) (*Task, *gorm.DB) {
	var getTask Task
	db := db.Where("ID = ?", Id).Find(&getTask)
	return &getTask, db
}

func DeleteTask(Id int64) Task {
	var task Task
	db.Where("ID = ?", Id).Delete(task)
	return task
}
