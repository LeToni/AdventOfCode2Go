package main

import (
	"bytes"
	"fmt"
)

var (
	password []byte
)

type Rule func() bool

func CheckPassword(rules ...Rule) bool {
	if len(rules) == 0 {
		return true
	}

	for _, rule := range rules {

		if !rule() {
			return false
		}
	}

	return true
}

func lettersIncreaseStraight() bool {

	for i := 0; i < len(password)-2; i++ {
		if password[i]+1 == password[i+1] && password[i+1]+1 == password[i+2] {
			return true
		}
	}

	return false
}

func lettersAreExcluded() bool {
	excludedLetters := "iol"
	return !bytes.ContainsAny(password, excludedLetters)
}

func twoPairsOfLetter() bool {
	pairs := 0
	for i := 0; i < len(password)-1; i++ {
		if password[i] == password[i+1] {
			pairs++
			// avoid overlapping pair detection
			i++
		}
	}

	return pairs >= 2
}

func PasswordIsValid() bool {
	return CheckPassword(twoPairsOfLetter,
		lettersAreExcluded,
		lettersIncreaseStraight,
	)
}

func incrementLastLetter() {
	for i := len(password) - 1; i > 0; i-- {
		if password[i] == 'z' {
			password[i] = 'a'
		} else {
			password[i] = password[i] + 1
			break
		}

	}

	password = []byte(password)
}

func main() {
	password = []byte("cqjxjnds")
	nextPassword()
	fmt.Println("1: New password according to rules: ", string(password))
	nextPassword()
	fmt.Println("2: New password according to rules: ", string(password))
}

func nextPassword() {
	validPassword := false
	for !validPassword {
		incrementLastLetter()
		validPassword = PasswordIsValid()
	}
}
