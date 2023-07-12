package remove

import (
	"context"
	"fmt"
	"os"

	"github.com/apex/log"
	longteaConfig "github.com/chungjung-d/longtea/config"
)

// RemoveContainer removes a container

func RemoveImage(ctx context.Context) {

	imageName := ctx.Value(longteaConfig.RemoveImageName).(string)
	imageDir := longteaConfig.GetImageDir()
	fullImagePath := fmt.Sprintf("%s/%s", imageDir, imageName)

	fmt.Println(fullImagePath)

	err := os.RemoveAll(fullImagePath)
	if err != nil {
		log.Fatal("Failed to remove image")
	}

}
