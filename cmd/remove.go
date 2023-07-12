/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"

	longteaConfig "github.com/chungjung-d/longtea/config"
	"github.com/spf13/cobra"

	"github.com/chungjung-d/longtea/feature/remove"
)

var removeContainerName string
var removeImageName string

// removeCmd represents the remove command
var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "remove the exist container or image",
	Long:  `remove the exist container or image`,
	Run: func(cmd *cobra.Command, args []string) {

		ctx := context.Background()

		if removeContainerName != "" {
			ctx = context.WithValue(ctx, longteaConfig.RemoveContainerName, removeContainerName)
			remove.RemoveContainer(ctx)

			fmt.Println("remove container")

		}

		if removeImageName != "" {

			ctx = context.WithValue(ctx, longteaConfig.RemoveImageName, removeImageName)
			remove.RemoveImage(ctx)

			fmt.Println("remove image")
		}
	},

	PreRunE: func(cmd *cobra.Command, args []string) error {

		if removeContainerName != "" && removeImageName != "" {
			return fmt.Errorf("required one flag: -c, --container or -i, --image")
		}

		if removeContainerName == "" && removeImageName == "" {
			return fmt.Errorf("required one flag: -c, --container or -i, --image")
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)

	removeCmd.Flags().StringVarP(&removeContainerName, "container", "c", "", "Container name to remove")
	removeCmd.Flags().StringVarP(&removeImageName, "image", "i", "", "Image name to remove")
}
