package pull

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/containers/image/copy"
	"github.com/containers/image/signature"
	"github.com/containers/image/transports"
	"github.com/containers/image/types"
)

func PullImage(imageDir string, imageName string) ([]byte, error) {
	os.Chdir(imageDir)
	ctx := context.Background()
	policyContext, err := getPolicyContext()
	if err != nil {
		log.Fatal("Failed to get policy context")
	}
	srcRef, err := ParseImageName(imageName, "docker")
	if err != nil {
		log.Fatal("Invalid image name")
	}
	destRef, err := ParseImageName(imageName, "oci")
	if err != nil {
		log.Fatal("Failed to set destination name")
	}
	return copy.Image(ctx, policyContext, destRef, srcRef, &copy.Options{})

}

func ParseImageName(imgName string, transportType string) (types.ImageReference, error) {
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

func getPolicyContext() (*signature.PolicyContext, error) {
	var policy *signature.Policy
	var err error
	policy = &signature.Policy{Default: []signature.PolicyRequirement{signature.NewPRInsecureAcceptAnything()}}
	if err != nil {
		return nil, err
	}
	return signature.NewPolicyContext(policy)
}
