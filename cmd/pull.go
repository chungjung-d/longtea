/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// pullCmd represents the pull command
var pullCmd = &cobra.Command{
	Use:   "pull",
	Short: "Pull the image on hub",
	Long: `Pull the image on hub`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("pull called")
	},
}

func init() {
	rootCmd.AddCommand(pullCmd)

}
