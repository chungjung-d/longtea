/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"

	longteaConfig "github.com/chungjung-d/longtea/config"
	"github.com/chungjung-d/longtea/feature/start"
	"github.com/spf13/cobra"
)

var startContainerName string

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "start the exist container",
	Long:  `start the exist container.`,
	Run: func(cmd *cobra.Command, args []string) {

		ctx := context.Background()

		containerDir := longteaConfig.GetContainerDir()
		containerDirPath := fmt.Sprintf("%s/%s", containerDir, startContainerName)

		ctx = context.WithValue(ctx, longteaConfig.ContainerDirPath, containerDirPath)

		start.StartContainer(ctx)

	},

	PreRunE: func(cmd *cobra.Command, args []string) error {

		if startContainerName == "" {

			return fmt.Errorf("required flag: -n, --name <container name>")
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(startCmd)

	startCmd.Flags().StringVarP(&startContainerName, "name", "n", "", "Container name to start")

}
