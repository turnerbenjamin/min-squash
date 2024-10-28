package helpers

import (
	"errors"
	"log"
	"os"
	"path/filepath"
	"regexp"

	"github.com/tdewolff/minify/v2"
	"github.com/tdewolff/minify/v2/css"
	"github.com/tdewolff/minify/v2/html"
	"github.com/tdewolff/minify/v2/js"
)

func Minify(srcPath string, tgtPath string) {
	if !IsModified(srcPath, tgtPath) {
		return
	}

	m := minify.New()
	m.AddFunc("text/css", css.Minify)
	m.AddFunc("text/html", html.Minify)
	m.AddFuncRegexp(regexp.MustCompile("^(application|text)/(x-)?(java|ecma)script$"), js.Minify)

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

	ext := filepath.Ext(srcPath)
	mediaType := ""

	switch ext {
	case ".html":
		mediaType = "text/html"
	case ".css":
		mediaType = "text/css"
	case ".js":
		mediaType = "text/javascript"
	default:
		log.Fatal(errors.New("invalid filetype"))
	}

	if err = m.Minify(mediaType, target, src); err != nil {
		log.Fatal(err)
	}
}
