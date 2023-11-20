package models

import (
	"GoTodoList/utils"
	"fmt"
)

type TaskModel struct {
	ID          uint `gorm:"primaryKey"`
	Name        string
	Description string
	Status      string
}

func InsertTask(taskObject *TaskModel) {
	database := utils.GetDB()
	database.Create(taskObject)

	fmt.Printf("Task created with ID: %v", taskObject.ID)
}

func SelectAllTask(tasks *[]TaskModel) {
	database := utils.GetDB()
	database.Find(tasks)
}

func UpdateTask(id uint) {
	database := utils.GetDB()
	database.Model(
		&TaskModel{},
	).Where("id = ?", id).Update("status", "DONE")

	fmt.Println("Task was successfully updated")
}

func DeleteTask(taskID uint) {
	database := utils.GetDB()
	database.Delete(&TaskModel{ID: taskID})

	fmt.Printf("Task with ID (%v) was deleted ", taskID)
}
