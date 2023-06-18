package config

import "fmt"

func GetImageDir() string {
	longteaDir := getenv("LONGTEA_DIR", "/var/lib/longtea")
	longteaImageDir := fmt.Sprintf("%s/images", longteaDir)
	return longteaImageDir
}
