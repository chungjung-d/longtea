package start

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"syscall"

	longteaConfig "github.com/chungjung-d/longtea/config"
)

func StartContainer(ctx context.Context) {

	containerDirPath := ctx.Value(longteaConfig.ContainerDirPath).(string)

	cmd := exec.Command("longteac", "create", "-c", containerDirPath)
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWPID | syscall.CLONE_NEWNS | syscall.CLONE_NEWIPC,
	}

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Start(); err != nil {
		fmt.Printf("Error running the /proc/self/exe - %s\n", err)
		os.Exit(1)
	}

	//나중에 DB연결 해서 컨테이너 정보 저장

	cmd.Wait()

}
