/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use: "config",
	Run: func(cmd *cobra.Command, args []string) {
		viper.WriteConfig()
	},
}

func init() {
	loadConfig()
	configCmd.Short = viper.GetString("cmd.config.short")
	configCmd.Long = viper.GetString("cmd.config.long")

	configCmd.Flags().StringP("data.dir", "d", "", "data directory")
	viper.BindPFlag("data.dir", configCmd.Flags().Lookup("data.dir"))

	rootCmd.AddCommand(configCmd)

}
