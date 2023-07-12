package remove

import (
	"context"
	"fmt"
	"os"

	"github.com/apex/log"
	longteaConfig "github.com/chungjung-d/longtea/config"
)

// RemoveContainer removes a container

func RemoveContainer(ctx context.Context) {

	containerName := ctx.Value(longteaConfig.RemoveContainerName).(string)
	containerDir := longteaConfig.GetContainerDir()
	fullContainerPath := fmt.Sprintf("%s/%s", containerDir, containerName)

	err := os.RemoveAll(fullContainerPath)
	if err != nil {
		log.Fatal("Failed to remove container")
	}

}
