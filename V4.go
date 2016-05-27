package isValidIP

import (
	"fmt"
	"regexp"
	"strings"
)

var (
	v4 *regexp.Regexp
)

func init() {
	var err error
	v4, err = regexp.Compile(`^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$`)
	if err != nil {
		fmt.Printf("isValidIP v4 initialize failed: %s", err)
	}
}

// V4 validates given IPv4
func V4(ip string) (bool, error) {
	ip = strings.Trim(ip, " ")

	if v4 == nil {
		return false, fmt.Errorf("Can't initialize validator")
	}

	if v4.MatchString(ip) {
		return true, nil
	}
	return false, nil
}
