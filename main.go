package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	result := genPass("abcd", "strong")
	fmt.Print("Password Result :", result)

}

func genPass(password string, level string) string {
	temp := ""
	passwords := strings.ToLower(password)
	special := []string{"!", "@", "#", "$", "&", "*"}
	number := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}

	type1 := strings.ToUpper(password[0:2])
	type2 := special[rand.Intn(len(special))]
	type3 := number[rand.Intn(len(number))]
	type4 := number[rand.Intn(len(number))]
	type5 := number[rand.Intn(len(number))]

	if len(passwords) < 6 {
		temp += "Password length minimum is 6 character"
	} else if level == "low" {
		temp += type1 + password + " Level: Low"
	} else if level == "medium" {
		temp += type1 + password + type3 + type4 + type5 + " Level: Medium"
	} else if level == "strong" {
		temp += type1 + password + type2 + type3 + type4 + type5 + " Level: Strong"
	} else {
		temp += "Something Wrong"
	}

	return temp
}
