/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"GoTodoList/models"

	"github.com/spf13/cobra"
)

var doneCmd = &cobra.Command{
	Use:   "done",
	Short: "Mark the task as completed",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		taskID, _ := cmd.Flags().GetUint("id")

		models.UpdateTask(taskID)
	},
}

func init() {
	rootCmd.AddCommand(doneCmd)

	doneCmd.PersistentFlags().UintP("id", "i", 0, "Task ID")
}
