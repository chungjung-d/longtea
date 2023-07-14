package image

import (
	"strings"
)

func ParseImageName(name string) (string, string) {
	split := strings.Split(name, ":")
	if len(split) != 2 {
		return split[0], ""
	}
	return split[0], split[1]
}
