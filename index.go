package golee

import (
	"strings"
)


func CensorEmail(email string) {
	var str_split = strings.Split(email, "@")
	censor_email := "xxxxxxx@" 

	result := censor_email + str_split[1]
	return result
}