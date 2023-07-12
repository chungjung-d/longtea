/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
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

		ctx := context.Background()

		ctx = context.WithValue(ctx, longteaConfig.CreateContainerName, createContainerName)
		ctx = context.WithValue(ctx, longteaConfig.CreateContainerOriginImageName, createImageName)

		create.CreateContainer(ctx)
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
