/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"
	"strings"

	longteaConfig "github.com/chungjung-d/longtea/config"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "longtea",
	Short: "Longtea is the container runtime",
	Long:  `Longtea is the container runtime. It just a sideproject to learn how the container works. You can use it just test or study`,
}

func Execute() {

	os.MkdirAll(longteaConfig.GetContainerDir(), os.ModePerm)
	os.MkdirAll(longteaConfig.GetImageDir(), os.ModePerm)

	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func splitContinaerName(name string) (string, string) {
	split := strings.Split(name, ":")
	if len(split) != 2 {
		return split[0], ""
	}
	return split[0], split[1]
}
