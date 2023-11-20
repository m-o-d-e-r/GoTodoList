/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"GoTodoList/models"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Insert a new task into TODO-list",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		taskName, _ := cmd.Flags().GetString("name")
		taskDesc, _ := cmd.Flags().GetString("desc")
		taskStatus, _ := cmd.Flags().GetString("status")

		if taskName == "" || taskDesc == "" || taskStatus == "" {
			cmd.Help()
			return
		}

		models.InsertTask(
			&models.TaskModel{
				Name:        taskName,
				Description: taskDesc,
				Status:      taskStatus,
			},
		)

	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	addCmd.PersistentFlags().StringP("name", "n", "", "Task name")
	addCmd.PersistentFlags().StringP("desc", "d", "", "Describe task")
	addCmd.PersistentFlags().StringP("status", "s", "NEW", "Task execution status")
}
