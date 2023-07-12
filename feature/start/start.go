package start

import (
	"context"
	"fmt"
	"os"
	"os/exec"

	longteaConfig "github.com/chungjung-d/longtea/config"
	"golang.org/x/sys/unix"
)

func StartContainer(ctx context.Context) {

	containerDirPath := ctx.Value(longteaConfig.ContainerDirPath).(string)

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
