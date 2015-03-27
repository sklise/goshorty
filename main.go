package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	home := os.Getenv("HOME")
	homeArr := strings.Split(home, "/")

	bytes, _ := ioutil.ReadAll(os.Stdin)
	pwd := string(bytes)
	// Remove the newline that comes out of `pwd`
	pwd = pwd[0:(len(pwd) - 1)]
	pwdArr := strings.Split(pwd, "/")

	comparisons := 0

	if len(pwdArr) >= len(homeArr) {
		for i := range homeArr {
			if homeArr[i] == pwdArr[i] {
				comparisons++
			}
		}
	}

	if comparisons == len(homeArr) {
		pwdArr = append([]string{"~"}, pwdArr[len(homeArr):]...)
	}

	output := []string{}

	lastPwdArr := len(pwdArr) - 1

	for i := 0; i < lastPwdArr; i++ {
		shortened := ""
		if len(pwdArr[i]) != 0 {
			shortened = pwdArr[i][:1]
		}

		output = append(output, shortened)
	}

	output = append(output, pwdArr[lastPwdArr])

	fmt.Print(strings.Join(output, "/"))
}
