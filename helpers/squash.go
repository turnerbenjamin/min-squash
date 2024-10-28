package helpers

import (
	"compress/gzip"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/andybalholm/brotli"
)

func CompressFile(srcPath string, destinationPath string) {
	if !IsModified(srcPath, destinationPath) {
		return
	}

	src, err := os.Open(srcPath)
	if err != nil {
		log.Fatal(err)
	}
	defer src.Close()

	destination, err := os.Create(destinationPath)
	if err != nil {
		log.Fatal(err)
	}
	defer destination.Close()

	if strings.HasSuffix(destinationPath, ".gz") {
		compressGzip(src, destination)
		return
	}

	if strings.HasSuffix(destinationPath, ".br") {
		compressBrotli(src, destination)
		return
	}

	fmt.Println("WARNING: destinationPath must end in gz or br")
}

func compressGzip(src *os.File, dWriter io.Writer) {
	gzipWriter, err := gzip.NewWriterLevel(dWriter, gzip.BestCompression)
	if err != nil {
		log.Fatal(err)
	}
	defer gzipWriter.Close()

	// Copy the contents of the original file to the gzip writer
	_, err = io.Copy(gzipWriter, src)
	if err != nil {
		log.Fatal(err)
	}

	gzipWriter.Flush()
}

func compressBrotli(src *os.File, dWriter io.Writer) {
	brotliWriter := brotli.NewWriterLevel(dWriter, brotli.BestCompression)
	defer brotliWriter.Close()

	// Copy the contents of the original file to the gzip writer
	_, err := io.Copy(brotliWriter, src)
	if err != nil {
		log.Fatal(err)
	}

	brotliWriter.Flush()
}
