/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	longteaConfig "github.com/chungjung-d/longtea/config"
	"github.com/chungjung-d/longtea/feature/create"
	"github.com/spf13/cobra"
)

var (
	createImageName     string
	createContainerName string
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "create the container with image",
	Long:  `create the container with image`,
	Run: func(cmd *cobra.Command, args []string) {

		imageDir := longteaConfig.GetImageDir()
		containerDir := longteaConfig.GetContainerDir()
		create.CreateContainer(imageDir, containerDir, createImageName, createContainerName)
	},

	PreRunE: func(cmd *cobra.Command, args []string) error {

		if createImageName == "" {
			return fmt.Errorf("required flag: -i, --image <image name>")
		}

		if createContainerName == "" {
			createContainerName = createImageName
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(createCmd)

	createCmd.Flags().StringVarP(&createImageName, "image", "i", "", "ImageName for create container")
	createCmd.Flags().StringVarP(&createContainerName, "name", "n", "", "ContainerName")
	createCmd.MarkFlagRequired("createImageName")
}
