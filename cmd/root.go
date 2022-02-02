package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "apachedist-resource",
	Short: "Implementation of a concourse resource that queries for updates on an apache dist",
	Long:  `Implementation of a concourse resource that queries for updates on an apache dist`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	rootCmd.AddCommand(NewInCmd().Command)
	rootCmd.AddCommand(NewCheckCmd().Command)

	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
