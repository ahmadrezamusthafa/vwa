package helper

import (
	"errors"
	"fmt"
	"regexp"
)

func NormalizeUID(uid *string) (string, error) {
	var regex, err = regexp.Compile("[^0-9]*")
	if err != nil {
		fmt.Println(err.Error())
	}

	if regex.MatchString(*uid) {
		*uid = regex.ReplaceAllString(*uid, "")
	}

	if *uid == "" {
		return *uid, errors.New("UID is invalid")
	}

	return *uid, nil
}
