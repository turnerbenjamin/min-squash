package helpers

import (
	"io"
	"log"
	"os"
)

func Copy(srcPath string, tgtPath string) {
	if !IsModified(srcPath, tgtPath) {
		return
	}
	src, err := os.Open(srcPath)
	if err != nil {
		log.Fatal(err)
	}
	defer src.Close()

	target, err := os.Create(tgtPath)
	if err != nil {
		log.Fatal(err)
	}
	defer target.Close()

	_, err = io.Copy(target, src)
	if err != nil {
		log.Fatal(err)
	}
}
