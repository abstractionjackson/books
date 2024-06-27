/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/abstractionjackson/books/library"
	"github.com/abstractionjackson/books/prompt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use: "update",
	PreRun: func(cmd *cobra.Command, args []string) {
		if title == "" {
			title = prompt.RunPromptTitle()
		}
		if status == "" {
			status = prompt.RunPromptStatus()
		}
		book = library.FindBookByTitle(title)
		// while books is nil, run not found prompt
		for book == nil {
			action := prompt.RunPromptNotFound()
			switch action {
			case "search again":
				title = prompt.RunPromptTitle()
				book = library.FindBookByTitle(title)
			case "add":
				// TODO add book
				fmt.Println("Feature in development: Pardon Our Dust")
				continue
			case "exit":
				os.Exit(0)
			}
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Updating %s to %s\n", title, status)
		book.UpdateStatus(status)
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
	updateCmd.Short = viper.GetString("cmd.update.short")
	updateCmd.Long = viper.GetString("cmd.update.long")
}
