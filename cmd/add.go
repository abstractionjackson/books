/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/abstractionjackson/books/library"
	"github.com/abstractionjackson/books/prompt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use: "add",
	PreRun: func(cmd *cobra.Command, args []string) {
		// if author, title, and status are not provided, run prompt
		if title == "" {
			title = prompt.RunPromptTitle()
		}
		if author == "" {
			author = prompt.RunPromptAuthor()
		}
		if status == "" {
			status = prompt.RunPromptStatus()
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		book = library.NewBook(title, author, status)
		book.SaveBookToTOML()
	},
}

func init() {
	loadConfig()
	addCmd.Short = viper.GetString("cmd.add.short")
	addCmd.Long = viper.GetString("cmd.add.long")
	rootCmd.AddCommand(addCmd)
}
