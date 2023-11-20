/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"GoTodoList/models"

	"github.com/rodaine/table"
	"github.com/spf13/cobra"
)

var showCmd = &cobra.Command{
	Use:   "show",
	Short: "Show all tasks in the form of a table",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		var tasks []models.TaskModel
		models.SelectAllTask(&tasks)

		tbl := table.New("ID", "NAME", "DESCRIPTION", "STATUS")

		for _, task := range tasks {
			tbl.AddRow(task.ID, task.Name, task.Description, task.Status)
		}

		tbl.Print()
	},
}

func init() {
	rootCmd.AddCommand(showCmd)
}
