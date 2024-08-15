package eancheckr

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var (
	containsAlpha   = regexp.MustCompile(`(?i)[a-z]+`).FindAllString
	replaceNonDigit = regexp.MustCompile(`\D`).ReplaceAllString
	stripZero       = func(in string) string { return strings.TrimLeft(in, "0") }
)

func VerifyEan(in string) (int, error) {

	if alphaChars := containsAlpha(in, 3); len(alphaChars) > 0 {
		return 0, fmt.Errorf("ean contains characters: %q", alphaChars)
	}

	checkValue := replaceNonDigit(in, "")

	if len(checkValue) < 8 {
		return 0, fmt.Errorf("ean must be at least 8 digits long: %q (%d) given", checkValue, len(checkValue))
	}

	if len(checkValue) > 14 {
		return 0, fmt.Errorf("ean must be at most 14 digits long: %q (%d) given", checkValue, len(checkValue))
	}

	if len(checkValue) < 14 {
		checkValue = fmt.Sprintf("%014s", checkValue)
	}

	var checkSum int
	for i := 0; i < 13; i++ {
		if i%2 == 0 {
			checkSum += int(checkValue[i]-'0') * 3
		} else {
			checkSum += int(checkValue[i] - '0')
		}
	}

	checkDigit := fmt.Sprintf("%v", (10-(checkSum%10))%10)

	correctEan := stripZero(checkValue[:13] + checkDigit)

	if checkDigit != string(checkValue[13]) {
		return 0, fmt.Errorf("ean check digit is wrong: %q given, expect %q", stripZero(checkValue), correctEan)
	}

	result, err := strconv.Atoi(correctEan)
	if err != nil {
		return 0, fmt.Errorf("failed to convert to int: %v", err)
	}

	return result, nil
}
