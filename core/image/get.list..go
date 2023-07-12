package image

import (
	"os"
	longteaConfig "github.com/chungjung-d/longtea/config"


	"github.com/apex/log"
)

func GetImageList() []string {

	imageDir := longteaConfig.GetImageDir()

	entries, err := os.ReadDir(imageDir)
	if err != nil {
		log.Fatal(err.Error())
	}

	var imageList []string
	for _, entry := range entries {
		if entry.IsDir() {
			imageList = append(imageList, entry.Name())
		}
	}

	return imageList
}
