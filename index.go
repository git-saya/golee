package main

import (
	"strings"
	"fmt"
)


func CensorEmail(email string) string {
	var str_split = strings.Split(email, "@")

	length := len(str_split[0])
	var censor_email string
	for i := 0; i < length; i++ {
		censor_email += "x" 
	}

	result := censor_email + "@" + str_split[1]
	return result
}

func main() {
	fmt.Println(CensorEmail("bale@gmail.com"))
}