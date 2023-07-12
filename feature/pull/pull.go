package pull

import (
	"context"
	"fmt"
	"os"

	"github.com/apex/log"

	longteaConfig "github.com/chungjung-d/longtea/config"
	longteaImage "github.com/chungjung-d/longtea/core/image"

	"github.com/containers/image/v5/copy"
	_ "github.com/containers/image/v5/docker"
	_ "github.com/containers/image/v5/oci/layout"
)

// TODO validate the exsit images
func PullImage(ctx context.Context) ([]byte, error) {

	imageName := ctx.Value(longteaConfig.PullImageName).(string)

	imageDir := longteaConfig.GetImageDir()

	os.Chdir(imageDir)

	name, tag := longteaImage.ParseImageName(imageName)

	imageName = fmt.Sprintf("%s:%s", name, tag)

	context := context.Background()
	policyContext, err := longteaImage.GetImagePolicy()
	if err != nil {
		log.Fatal("Failed to get policy context")
	}
	srcRef, err := longteaImage.GetImageReference(imageName, "docker")
	if err != nil {
		log.Fatal("Invalid image name")
	}
	destRef, err := longteaImage.GetImageReference(imageName, "oci")
	if err != nil {
		log.Fatal("Failed to set destination name")
	}
	return copy.Image(context, policyContext, destRef, srcRef, &copy.Options{})

}
