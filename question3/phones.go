package question3

import (
	"errors"
	"regexp"
	"strings"
)

const NANP_REGEX = `(\+1 |1 )?\(?[2-9][0-9]{2}\)?[.\- ]?[2-9][0-9]{2}[.\- ]?\d{4}`

func Number(phoneNumber string) (string, error) {
	result := ""
	match, err := regexp.Match(NANP_REGEX, []byte(phoneNumber))
	if err != nil {
		return result, err
	}
	if match {

		result = strings.Map(func(r rune) rune {
			switch r {
			case '(', ')', '-', '+', ' ', '.':
				return -1
			default:
				return r
			}
		}, phoneNumber)

		if len(result) > 10 {
			return result[1:], err
		}

		return result, nil
	}
	return "", errors.New("not a valid NANP phone number")
}
