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

// infoCmd represents the info command
var infoCmd = &cobra.Command{
	Use: "info",
	PreRun: func(cmd *cobra.Command, args []string) {
		// if no title, prompt
		if title == "" {
			title = prompt.RunPromptTitle()
		}
		for book == nil {
			book = library.FindBookByTitle(title)
			if book == nil {
				action := prompt.RunPromptNotFound()
				switch action {
				case "search again":
					title = prompt.RunPromptTitle()
				case "add":
					fmt.Println("Work in progress")
				case "exit":
					os.Exit(0)
				}
			}
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Title:", book.Title)
		fmt.Println("Author:", book.Author)
		status, date := book.GetCurrentStatusAndDate()
		daysAtStatus := int(date.Sub(date).Hours() / 24)
		fmt.Println("Status:", status)
		fmt.Printf("Date: %s (%d days at %s)\n", date.Format("January 2, 2006"), daysAtStatus, status)
	},
}

func init() {
	rootCmd.AddCommand(infoCmd)
	infoCmd.Short = viper.GetString(("cmd.info.short"))
	infoCmd.Long = viper.GetString(("cmd.info.long"))
}
