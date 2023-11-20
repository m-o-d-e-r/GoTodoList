/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"GoTodoList/models"

	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Remove task from the TODO-list",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		taskID, _ := cmd.Flags().GetUint("id")

		models.DeleteTask(taskID)
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)

	deleteCmd.PersistentFlags().UintP("id", "i", 0, "Task ID to remove")
}
