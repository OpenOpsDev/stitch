package templates

import (
	"errors"
	"strings"

	"github.com/aymerick/raymond"
	"github.com/roger-king/stitch/utils"
)

type HandlebarTemplate struct {
	// Raw - is the raw string from a file
	Raw         string
	Destination string
}

func (h HandlebarTemplate) Render(vars map[string]string) (string, error) {
	tpl, err := raymond.Parse(h.Raw)
	if err != nil {
		panic(err)
	}

	return tpl.Exec(vars)
}

func (h HandlebarTemplate) Save(vars map[string]string) error {
	result, err := h.Render(vars)

	if err != nil {
		return err
	}

	// Write to file
	nilGeneratedContent := len(result) == 0
	if nilGeneratedContent {
		return errors.New("no generated output available")
	}

	splitPath := strings.Split(h.Destination, "/")
	fileNameWithExtension := splitPath[len(splitPath)-1]
	fileName := strings.Split(fileNameWithExtension, ".hbs")[0]
	splitPath = splitPath[0 : len(splitPath)-1]
	fileDest := strings.Join(splitPath, "/")

	file := utils.File{
		Name: fileName,
		Data: result,
		Dest: fileDest,
	}
	return file.Write()
}
