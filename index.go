package golee

import (
	"strings"
)

func CensorEmail(email string) string {
	var str_split = strings.Split(email, "@")
	initial := str_split[0][0:1] //ambil huruf pertama

	length := len(str_split[0][1:]) //hitung jumlah string mulai dari index 1: kecuali huruf pertama

	var censor_email string
	for i := 1; i < length; i++ {
		censor_email += "x" 
	}

	// gabungkan semua variabel
	result := initial + censor_email + "@" + str_split[1]
	return result
}