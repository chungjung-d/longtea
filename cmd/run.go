/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"

	longteaConfig "github.com/chungjung-d/longtea/config"
	"github.com/chungjung-d/longtea/feature/run"
	"github.com/spf13/cobra"
)

var runContainerName string
var runContainerImageName string

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "run the new container",
	Long:  `run the new container`,
	Run: func(cmd *cobra.Command, args []string) {

		ctx := context.Background()

		containerDir := longteaConfig.GetContainerDir()
		imagerDir := longteaConfig.GetImageDir()
		containerDirPath := fmt.Sprintf("%s/%s", containerDir, runContainerName)
		imageDirPath := fmt.Sprintf("%s/%s", imagerDir, runContainerImageName)

		ctx = context.WithValue(ctx, longteaConfig.ContainerDirPath, containerDirPath)
		ctx = context.WithValue(ctx, longteaConfig.ImageDirPath, imageDirPath)

		run.RunContainer(ctx)
	},

	PreRunE: func(cmd *cobra.Command, args []string) error {

		if runContainerImageName == "" {

			return fmt.Errorf("required flag: -i, --image <image name>")
		}

		if runContainerName == "" {

			runContainerName = runContainerImageName

		}

		containerName, containerTags := splitContinaerName(runContainerName)
		if containerTags != "" {
			runContainerName = fmt.Sprintf("%s_%s", containerName, containerTags)
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(runCmd)

	runCmd.Flags().StringVarP(&runContainerImageName, "image", "i", "", "ImageName for create container")
	runCmd.Flags().StringVarP(&runContainerName, "name", "n", "", "Container name to start")

}
