package create

import (
	"fmt"

	"github.com/apex/log"

	longteaContainer "github.com/chungjung-d/longtea/core/container"
	longteaImage "github.com/chungjung-d/longtea/core/image"
)

func CreateContainer(imageDir string, containerDir string, imageName string, containerName string) {
	imageNamePart, imageTagPart := longteaImage.ParseImageName(imageName)

	containerExist := longteaContainer.ValidateExistContainer(containerDir, containerName)
	if containerExist {
		log.Fatal("Container already exist")
	}

	err := longteaContainer.CreateContainer(imageDir, containerDir, containerName, imageNamePart, imageTagPart)
	if err != nil {
		log.Fatal(fmt.Sprintf("Failed to create container. Error received: %s", err.Error()))
	}

}
