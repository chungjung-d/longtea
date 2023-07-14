package image

import (
	"fmt"
	"os"
	"path/filepath"

	longteaConfig "github.com/chungjung-d/longtea/config"
)

func GetImageList() []string {

	imageDir := longteaConfig.GetImageDir()

	entries, err := os.ReadDir(imageDir)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	var imageList []string
	for _, entry := range entries {
		if entry.IsDir() {
			imageName := entry.Name()
			tags := longteaConfig.GetImageTag(filepath.Join(imageDir, imageName))
			for _, tag := range tags {
				imageList = append(imageList, fmt.Sprintf("%s:%s", imageName, tag))
			}
		}
	}

	return imageList
}
