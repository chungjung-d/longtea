package container

import (
	"os"

	"github.com/apex/log"
)

func GetContainerList(containerDir string) []string {
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

func ValidateExistContainer(containerDir string, containerName string) bool {
	containerList := GetContainerList(containerDir)

	for _, container := range containerList {
		if container == containerName {
			return true
		}
	}

	return false

}
