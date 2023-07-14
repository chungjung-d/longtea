package run

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"golang.org/x/sys/unix"

	"github.com/apex/log"

	longteaConfig "github.com/chungjung-d/longtea/config"
	longteaContainer "github.com/chungjung-d/longtea/core/container"
	longteaImage "github.com/chungjung-d/longtea/core/image"
)

func RunContainer(ctx context.Context) {
	containerDirPath := ctx.Value(longteaConfig.ContainerDirPath).(string)
	ImageDirPath := ctx.Value(longteaConfig.ImageDirPath).(string)

	_, err := os.Stat(containerDirPath)
	if err == nil {
		fmt.Println("The container is exist. If you want to start it, use start command ")
		fmt.Print(" Or continue, all the existing data will be erased. Shall we contineu? [y/n]: ")

		var answer string
		fmt.Scanln(&answer)
		if answer != "y" {
			log.Fatal("Stop the process by user")
		}

		os.RemoveAll(containerDirPath)
	}

	imageName := strings.Split(ImageDirPath, "/")[len(strings.Split(ImageDirPath, "/"))-1]
	imageNamePart, imageTagPart := longteaImage.ParseImageName(imageName)
	if imageTagPart == "" {
		imageTagPart = "latest"
	}

	containerName := strings.Split(containerDirPath, "/")[len(strings.Split(containerDirPath, "/"))-1]

	// TODO Refactor this part
	err = longteaContainer.CreateContainer(longteaConfig.GetImageDir(), longteaConfig.GetContainerDir(), containerName, imageNamePart, imageTagPart)
	if err != nil {
		log.Fatal(fmt.Sprintf("Failed to create container. Error received: %s", err.Error()))
	}

	cmd := exec.Command("longteac", "create", "-c", containerDirPath)
	cmd.SysProcAttr = &unix.SysProcAttr{
		Cloneflags: unix.CLONE_NEWPID | unix.CLONE_NEWNS | unix.CLONE_NEWIPC,
	}

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Start(); err != nil {
		fmt.Printf("Error running the /proc/self/exe - %s\n", err)
		os.Exit(1)
	}

	fmt.Println("Process ID is:", cmd.Process.Pid)

	cmd.Wait()

}
