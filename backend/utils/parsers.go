package utils

import (
	"fmt"
	"os"
)

func IsValidDirectory(path string) (bool, error) {
	info, err := os.Stat(path)
	fmt.Println(path)
	if err != nil {
		return false, err
	}

	return info.IsDir(), err
}
