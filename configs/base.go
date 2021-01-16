package configs

import (
	"os"

	"gopkg.in/yaml.v2"
)

func writeToFile(dest, result string) error {
	file, err := os.Create(dest)

	if err != nil {
		return err
	}

	defer file.Close()

	_, err = file.WriteString(result)

	if err != nil {
		return err
	}

	return nil
}

// Render -
func Render(dest string, s interface{}) error {
	results, err := yaml.Marshal(s)

	if err != nil {
		return err
	}

	err = writeToFile(dest, string(results))

	if err != nil {
		return err
	}

	return nil
}
