package golee

import (
	"strings"
)

func CensorEmail(email string) string {
	var str_split = strings.Split(email, "@")
	initial := str_split[0][0:1] //ambil huruf pertama

	length := len(str_split[0][1:]) //hitung jumlah string mulai dari index 1: kecuali huruf pertama

	var censor_email string
	for i := 0; i < length; i++ {
		censor_email += "x" 
	}

	// gabungkan semua variabel
	result := initial + censor_email + "@" + str_split[1]
	return result
}

func CensorPhoneNumber(phone string) string {
	length  := len(phone)
	initial := phone[0:4]

	// slice
	lastDigit := phone[length-4:] // ambil karakter terakhir 4 digit.
	var censor_number string
	for i := 4; i < length-4; i++ {
		censor_number += "x" 
	}

	result := initial + censor_number + lastDigit
	return result;
}