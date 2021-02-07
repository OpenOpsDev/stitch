package configs

import (
	"github.com/roger-king/stitch/utils"
	"gopkg.in/yaml.v2"
)

// Render -
func Render(dest string, s interface{}) error {
	results, err := yaml.Marshal(s)

	if err != nil {
		return err
	}
	file := utils.File{
		Data: string(results),
		Dest: dest,
	}
	err = file.Write()

	if err != nil {
		return err
	}

	return nil
}
