/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/apex/log"
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

		if createImageName == "" {
			log.Fatal("image is required")
		}

		if createContainerName == "" {
			createContainerName = createImageName
		}

		imageDir := longteaConfig.GetImageDir()
		containerDir := longteaConfig.GetContainerDir()
		create.CreateContainer(imageDir, containerDir, createImageName, createContainerName)
	},

	
}

func init() {
	rootCmd.AddCommand(createCmd)

	createCmd.Flags().StringVarP(&createImageName, "image", "i", "", "ImageName for create container")
	createCmd.Flags().StringVarP(&createContainerName, "name", "n", "", "ContainerName")
	createCmd.MarkFlagRequired("createImageName")
}
