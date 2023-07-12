package list

import (
	"fmt"

	longteaImage "github.com/chungjung-d/longtea/core/image"
)

func ListImages() {

	imageList := longteaImage.GetImageList()

	for _, image := range imageList {
		fmt.Println(image)
	}

}
