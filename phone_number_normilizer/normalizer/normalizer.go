package normalizer

import (
	"fmt"
	"regexp"
)

func phoneNumber(phones ...string) ([]string, error) {
	if phones == nil || len(phones) == 0 {
		return nil, fmt.Errorf("'Phones' param can't be empty")
	}
	var output = make([]string, 0)
	for _, phone := range phones {
		rg := regexp.MustCompile(`\D`)
		replaced := rg.ReplaceAllLiteralString(phone, "")
		output = append(output, string(replaced))
	}
	return output, nil
}
