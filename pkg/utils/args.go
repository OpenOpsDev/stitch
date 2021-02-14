package utils

import "errors"

func CheckRequiredNumOfArgs(args []string, numRequired int) (bool, error) {
	if len(args) < numRequired {
		return false, errors.New("not enough arguments")
	}

	if len(args) > numRequired {
		return false, errors.New("too many arguments")
	}

	return true, nil
}
