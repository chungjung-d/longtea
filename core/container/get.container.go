package container

import (
	"os"

	longteaConfig "github.com/chungjung-d/longtea/config"
	

	"github.com/apex/log"
)

func GetContainerList() []string {

	containerDir := longteaConfig.GetContainerDir()

	entries, err := os.ReadDir(containerDir)
	if err != nil {
		log.Fatal(err.Error())
	}

	var containerList []string
	for _, entry := range entries {
		if entry.IsDir() {
			containerList = append(containerList, entry.Name())
		}
	}

	return containerList
}

func ValidateExistContainer(containerName string) bool {
	containerList := GetContainerList()

	for _, container := range containerList {
		if container == containerName {
			return true
		}
	}

	return false

}
