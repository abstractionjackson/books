/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/abstractionjackson/books/library"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var title, author, status string
var book *library.Book

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use: "books",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		dataDir := viper.GetString("data.dir")
		if _, err := os.Stat(dataDir); os.IsNotExist(err) {
			os.Mkdir(dataDir, 0755)
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("root called")
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func loadConfig() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.SetConfigType("toml")

	viper.SetDefault("data.dir", "data")
	viper.ReadInConfig()
}

func init() {
	loadConfig()

	rootCmd.Short = viper.GetString("cmd.root.short")
	rootCmd.Long = viper.GetString("cmd.root.long")

	rootCmd.PersistentFlags().StringVarP(&title, "title", "t", "", "book title")
	rootCmd.PersistentFlags().StringVarP(&author, "author", "a", "", "book author")
	rootCmd.PersistentFlags().StringVarP(&status, "status", "s", "", "book status")
}
