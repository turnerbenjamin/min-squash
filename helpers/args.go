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

func GetArgs() (*Args, error) {
	inputDirFlag := ""
	outputDirFlag := ""
	fileTypesFlag := ""

	flag.StringVar(&inputDirFlag, "sourcedir", inputDirFlag, "source directory")
	flag.StringVar(&inputDirFlag, "sd", inputDirFlag, "source directory")

	flag.StringVar(&outputDirFlag, "targetdir", outputDirFlag, "target directory")
	flag.StringVar(&outputDirFlag, "td", outputDirFlag, "target directory")

	flag.StringVar(&fileTypesFlag, "filetypes", fileTypesFlag, "filetypes to min-squash")
	flag.StringVar(&fileTypesFlag, "ft", fileTypesFlag, "filetypes to min-squash")

	flag.Parse()

	if inputDirFlag == "" {
		return nil, errors.New("input flag must be set")
	}

	if outputDirFlag == "" {
		return nil, errors.New("output flag must be set")
	}

	if fileTypesFlag == "" {
		return nil, errors.New("filetypes flag must be set")
	}

	if !strings.HasSuffix(inputDirFlag, "/") {
		inputDirFlag = inputDirFlag + "/"
	}

	if !strings.HasSuffix(outputDirFlag, "/") {
		outputDirFlag = outputDirFlag + "/"
	}

	filetypes := strings.Split(fileTypesFlag, ",")
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
		InputDir:  inputDirFlag,
		OutputDir: outputDirFlag,
		Filetypes: filetypes,
	}

	if userArgs.InputDir == userArgs.OutputDir {
		em := "ouput directory cannot equal input directory"
		return nil, errors.New(em)
	}

	return &userArgs, nil
}
