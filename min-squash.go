package minsquash

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/turnerbenjamin/minsquash/helpers"
)

func MinSquash(inputDir string, outputDir string, filetypes []string) {
	args, err := helpers.ParseArgs(inputDir, outputDir, filetypes)
	if err != nil {
		log.Fatal(err)
	}

	dir := helpers.GetDirFs(args.InputDir)
	files, err := helpers.GetFilesFromDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	for _, path := range files {
		srcPath := fmt.Sprintf("%s%s", args.InputDir, path)
		targetPath := fmt.Sprintf("%s%s", args.OutputDir, path)

		scaffoldDir(targetPath)

		if doSkip(path, args.Filetypes) {
			helpers.Copy(srcPath, targetPath)
			continue
		}

		helpers.Minify(srcPath, targetPath)

		// Compress GZip
		gzipTargetPath := fmt.Sprintf("%s.gz", targetPath)
		helpers.CompressFile(srcPath, gzipTargetPath)

		// Compress Brotli
		brTargetPath := fmt.Sprintf("%s.br", targetPath)
		helpers.CompressFile(srcPath, brTargetPath)

	}

	fmt.Println("min-squash successful")
}

func doSkip(path string, filetypes []string) bool {
	for _, ft := range filetypes {
		if strings.HasSuffix(path, ft) {
			return false
		}
	}
	return true
}

func scaffoldDir(targetPath string) {
	dirToBuild := "/"
	if strings.HasPrefix(targetPath, "./") {
		dirToBuild = "./"
	}
	for _, pathElement := range strings.Split(filepath.Dir(targetPath), "/") {
		dirToBuild = dirToBuild + pathElement + "/"
		mkDir(dirToBuild)
	}
}

func mkDir(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		if err = os.Mkdir(path, 0750); err != nil {
			log.Fatal(err)
		}
	}
}
