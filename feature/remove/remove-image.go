package remove

import (
	"context"
	"fmt"
	"os"

	longteaConfig "github.com/chungjung-d/longtea/config"
	longteaImage "github.com/chungjung-d/longtea/core/image"
)

// RemoveContainer removes a container

func RemoveImage(ctx context.Context) {

	imageName := ctx.Value(longteaConfig.RemoveImageName).(string)
	imageDir := longteaConfig.GetImageDir()

	imageNamePart, imageTagPart := longteaImage.ParseImageName(imageName)
	if imageTagPart == "" {
		os.Remove(fmt.Sprintf("%s/%s", imageDir, imageNamePart))
		return
	}

	fullImagePath := fmt.Sprintf("%s/%s", imageDir, imageNamePart)

	fmt.Println(fullImagePath)

	longteaConfig.RemoveImage(fullImagePath, imageTagPart)

}
