package config

import "fmt"

func GetContainerDir() string {
	longteaDir := getenv("LONGTEA_DIR", "/var/lib/longtea")
	longteaContainerDir := fmt.Sprintf("%s/containers", longteaDir)
	return longteaContainerDir
}
