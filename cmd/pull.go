/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/apex/log"

	longteaConfig "github.com/chungjung-d/longtea/config"
	"github.com/chungjung-d/longtea/feature/pull"

	"github.com/spf13/cobra"
)

var pullImageName string

// pullCmd represents the pull command
var pullCmd = &cobra.Command{
	Use:   "pull",
	Short: "Pull the image on hub",
	Long:  `Pull the image on hub`,
	Run: func(cmd *cobra.Command, args []string) {
		if pullImageName == "" {
			log.Fatal("pullImage is required")
		}

		imageDir := longteaConfig.GetImageDir()

		_, err := pull.PullImage(imageDir, pullImageName)
		if err != nil {
			log.Fatal(fmt.Sprintf("Failed to pull image. Error received: %s", err.Error()))
		}
	},
}

func init() {
	rootCmd.AddCommand(pullCmd)

	pullCmd.Flags().StringVarP(&pullImageName, "pullImageName", "i", "", "DockerImage Name to pull")
	pullCmd.MarkFlagRequired("pullImage")
}
