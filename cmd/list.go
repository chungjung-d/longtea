/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/chungjung-d/longtea/feature/list"
	"github.com/spf13/cobra"
)

var listImages bool
var listContainers bool

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list the containers or images",
	Long:  `list the containers or images`,
	Run: func(cmd *cobra.Command, args []string) {

		if listImages {

			fmt.Println("============== list images ==============")
			list.ListImages()
			fmt.Println()
		}

		if listContainers {
			
			fmt.Println("============== list containers ==============")
			list.ContinerImages()
			fmt.Println()
		}

	},

	PreRunE: func(cmd *cobra.Command, args []string) error {

		if !listImages && !listContainers {
			return fmt.Errorf("required flag: -i, --images or -c, --containers or both")
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	listCmd.Flags().BoolVarP(&listImages, "images", "i", false, "List images")
	listCmd.Flags().BoolVarP(&listContainers, "containers", "c", false, "List containers")
}
