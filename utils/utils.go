package utils

import (
	"net"
	"regexp"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
)

type email string

var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
var phoneRegex = regexp.MustCompile(`1?\W*([2-9][0-8][0-9])\W*([2-9][0-9]{2})\W*([0-9]{4})(\se?x?t?(\d*))?`)

/*
IsEmailValid :-
Description [Function checks if the email provided passes the required structure and length test. It also checks the domain has a valid MX record.]
@params [email string]
@return [true/false bool]
*/
func IsEmailValid(e string) bool {
	if len(e) < 3 && len(e) > 254 {
		logrus.Error("Util Function | IsEmailValid ", "Length issue")
		return false
	}
	if !emailRegex.MatchString(e) {
		logrus.Error("Util Function | IsEmailValid ", "Regex issue")
		return false
	}
	parts := strings.Split(e, "@")
	mx, err := net.LookupMX(parts[1])
	if err != nil || len(mx) == 0 {
		logrus.Error("Util Function | IsEmailValid ", err)
		return false
	}
	return true
}

/*
IsPhoneNoValid :-
Description [Function checks if the phone is valid or not]
@params [phoneNo string]
@return [true/false bool]
*/
func IsPhoneNoValid(e int) bool {
	phone := strconv.Itoa(e)
	if len(phone) != 10 {
		return false
	}

	return true
}
