package helpers

import (
	"errors"
	"flag"
	"strings"
)

type Args struct {
	InputDir  string
	OutputDir string
	Filetypes []string
}

func ParseArgs(inputDir string, outputDir string, filetypes []string) (*Args, error) {
	if inputDir == "" {
		return nil, errors.New("input dir must be set")
	}

	if outputDir == "" {
		return nil, errors.New("output dir must be set")
	}

	if filetypes == nil || len(filetypes) == 0 {
		return nil, errors.New("filetypes must be set")
	}

	if !strings.HasSuffix(inputDir, "/") {
		inputDir = inputDir + "/"
	}

	if !strings.HasSuffix(outputDir, "/") {
		outputDir = outputDir + "/"
	}

	for i, ft := range filetypes {
		formatted := strings.Trim(ft, " ")
		if strings.Contains(formatted, " ") {
			return nil, errors.New("filetypes cannot contain whitespace")
		}

		if !strings.HasPrefix(formatted, ".") {
			formatted = "." + formatted
		}

		filetypes[i] = formatted
	}
	userArgs := Args{
		InputDir:  inputDir,
		OutputDir: outputDir,
		Filetypes: filetypes,
	}

	if userArgs.InputDir == userArgs.OutputDir {
		em := "ouput directory cannot equal input directory"
		return nil, errors.New(em)
	}

	return &userArgs, nil
}
