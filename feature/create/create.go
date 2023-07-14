package create

import (
	"context"
	"fmt"

	"github.com/apex/log"

	longteaConfig "github.com/chungjung-d/longtea/config"
	longteaContainer "github.com/chungjung-d/longtea/core/container"
	longteaImage "github.com/chungjung-d/longtea/core/image"
)

// imageDir string, containerDir string, imageName string, containerName string
func CreateContainer(ctx context.Context) {

	imageName := ctx.Value(longteaConfig.CreateContainerOriginImageName).(string)
	containerName := ctx.Value(longteaConfig.CreateContainerName).(string)

	imageDir := longteaConfig.GetImageDir()
	containerDir := longteaConfig.GetContainerDir()

	imageNamePart, imageTagPart := longteaImage.ParseImageName(imageName)
	if imageTagPart == "" {
		imageTagPart = "latest"
	}

	fmt.Println(containerName)

	containerExist := longteaContainer.ValidateExistContainer(containerName)
	if containerExist {
		log.Fatal("Container already exist")
	}

	err := longteaContainer.CreateContainer(imageDir, containerDir, containerName, imageNamePart, imageTagPart)
	if err != nil {
		log.Fatal(fmt.Sprintf("Failed to create container. Error received: %s", err.Error()))
	}

}
