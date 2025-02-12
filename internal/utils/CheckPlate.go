package utils

import (
	"regexp"
)

const regexCheckPlate = `^((?:[A-Z]{2}-\d{3}-[A-Z]{2})|(?:\d{1,4}[ ][A-Z]{1,3}[ ]\d{2}))$`

// CheckPlate checks if the plate is valid
func CheckPlate(plate string) bool {
	return regexp.MustCompile(regexCheckPlate).MatchString(plate)
}
