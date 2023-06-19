package container

import (
	"fmt"
	"os"

	"github.com/opencontainers/umoci"
	"github.com/opencontainers/umoci/oci/cas/dir"
	"github.com/opencontainers/umoci/oci/casext"
	"github.com/opencontainers/umoci/oci/layer"
	"github.com/simonz05/util/log"
)

func CreateContainer(imageDir string, containerDir string, containerName string, imageName string, imageTag string) error {

	os.Chdir(imageDir)

	engine, err := dir.Open(imageName)

	if err != nil {
		log.Fatal(err)
	}

	engineExt := casext.NewEngine(engine)

	var unpackOptions layer.UnpackOptions
	var meta umoci.Meta

	meta.Version = umoci.MetaVersion
	meta.MapOptions.Rootless = true
	unpackOptions = layer.UnpackOptions{
		MapOptions: meta.MapOptions,
	}

	fullContainerPath := fmt.Sprintf("%s/%s", containerDir, containerName)
	return umoci.Unpack(engineExt, imageTag, fullContainerPath, unpackOptions)
}
