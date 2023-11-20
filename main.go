/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"GoTodoList/cmd"
	"GoTodoList/models"
	"GoTodoList/utils"
)

func main() {
	utils.CreateProgramFolder()
	utils.InitDB()

	db := utils.GetDB()
	db.AutoMigrate(&models.TaskModel{})

	cmd.Execute()
}
