package list

import (
	"fmt"

	longteaContainer "github.com/chungjung-d/longtea/core/container"
)

func ContinerImages() {

	containerList := longteaContainer.GetContainerList()

	for _, container := range containerList {
		fmt.Println(container)
	}

}
