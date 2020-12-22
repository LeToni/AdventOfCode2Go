package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"io/ioutil"
	"regexp"
	"strconv"
)

var (
	number    int
	secretKey string
)

func main() {

	input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	regex := regexp.MustCompile(`\A000000`)

	for !regex.MatchString(secretKey) {
		number = number + 1
		hash := md5.New()
		io.WriteString(hash, string(input)+strconv.Itoa(number))
		secretKey = fmt.Sprintf("%x", hash.Sum(nil))
	}

	fmt.Println(number)
}
