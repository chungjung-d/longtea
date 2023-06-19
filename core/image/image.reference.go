package image

import (
	"fmt"
	"strings"

	"github.com/apex/log"
	_ "github.com/containers/image/v5/docker"
	_ "github.com/containers/image/v5/oci/layout"
	"github.com/containers/image/v5/transports"
	"github.com/containers/image/v5/types"
)

func GetImageReference(imgName string, transportType string) (types.ImageReference, error) {
	fmt.Println(transportType)
	transport := transports.Get(transportType)
	if transport == nil {
		log.Fatal("Failed to get image transport type.")
	}
	imgNameSplit := strings.Split(imgName, "/")
	imgName = imgNameSplit[len(imgNameSplit)-1]
	if transportType == "docker" {
		imgName = fmt.Sprintf("//%s", imgName)
	}
	return transport.ParseReference(imgName)
}
