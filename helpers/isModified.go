package helpers

import (
	"log"
	"os"
)

func IsModified(srcPath string, destinationPath string) bool {
	srcInfo, err := os.Stat(srcPath)
	if err != nil {
		log.Fatal(err)
	}

	destinationInfo, err := os.Stat(destinationPath)
	if err != nil {
		return true
	}

	return srcInfo.ModTime().After(destinationInfo.ModTime())
}
