package main

import (
	"fmt"
	"os"
	"regexp"
)

func main() {
	var Reset = "\033[0m"
	var Red = "\033[31m"
	var Green = "\033[32m"

	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s%v%s\n", Red, err, Reset)
		os.Exit(1)
	}

	fmt.Fprintf(os.Stdout, "%svalid%s\n", Green, Reset)
}

func run() error {
	if len(os.Args) != 2 {
		return fmt.Errorf("usage: %s <ean>", os.Args[0])
	}

	return verifyEan(os.Args[1])
}

var containsAlpha = regexp.MustCompile(`(?i)[a-z]+`).FindAllString
var replaceNonDigit = regexp.MustCompile(`\D`).ReplaceAllString

func verifyEan(in string) error {

	if alphaChars := containsAlpha(in, 1); len(alphaChars) > 0 {
		return fmt.Errorf("ean contains alpha characters: %q", alphaChars[0])
	}

	checkValue := replaceNonDigit(in, "")

	if len(checkValue) < 8 {
		return fmt.Errorf("ean must be at least 8 digits long: %q given", checkValue)
	}

	if len(checkValue) > 14 {
		return fmt.Errorf("ean must be at most 14 digits long: %q given", checkValue)
	}

	checkValue = fmt.Sprintf("%014s", checkValue)

	var total int

	for i := 0; i < len(checkValue)-1; i++ {
		if i%2 == 0 {
			total += int(checkValue[i]-'0') * 3
		} else {
			total += int(checkValue[i] - '0')
		}
	}

	checkSum := fmt.Sprintf("%v", (10-(total%10))%10)

	correctEan := checkValue[:13] + checkSum

	if checkSum != string(checkValue[13]) {
		return fmt.Errorf("ean check digit is wrong: %q given, expect %q", checkValue, correctEan)
	}

	return nil
}
