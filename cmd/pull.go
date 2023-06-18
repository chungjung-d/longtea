/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"

	longteaConfig "github.com/chungjung-d/longtea/config"
	"github.com/chungjung-d/longtea/src/feature/pull"

	"github.com/spf13/cobra"
)

var imageName string

// pullCmd represents the pull command
var pullCmd = &cobra.Command{
	Use:   "pull",
	Short: "Pull the image on hub",
	Long:  `Pull the image on hub`,
	Run: func(cmd *cobra.Command, args []string) {
		if imageName == "" {
			log.Fatal("image is required")
		}

		imageDir := longteaConfig.GetImageDir()

		_, err := pull.PullImage(imageDir, imageName)
		if err != nil {
			log.Fatal(fmt.Sprintf("Failed to pull image. Error received: %s", err.Error()))
		}
	},
}

func init() {
	rootCmd.AddCommand(pullCmd)

	pullCmd.Flags().StringVarP(&imageName, "image", "i", "", "DockerImage Name to pull")
	pullCmd.MarkFlagRequired("image")
}
