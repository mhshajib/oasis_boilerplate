package utils

import (
	"fmt"
	"strings"
)

func ConvertPhoneToStandardFormat(mobile string) (string, error) {
	// Remove all leading "+" or "00" prefixes
	mobile = strings.TrimPrefix(mobile, "+")
	mobile = strings.TrimPrefix(mobile, "00")

	// With the "880" prefix, it should be 13 characters long
	if strings.HasPrefix(mobile, "880") {
		if len(mobile) != 13 {
			return "", fmt.Errorf("Invalid mobile number length")
		}
		return mobile, nil
	}

	// With only the "0" prefix, it should be 11 characters long
	if strings.HasPrefix(mobile, "0") {
		if len(mobile) != 11 || mobile[2] < '3' || mobile[2] > '9' {
			return "", fmt.Errorf("Invalid mobile number format")
		}
		return "880" + mobile[1:], nil
	}

	// Without the "880" or "0" prefix, it should be 10 characters long
	if len(mobile) == 10 && mobile[0] == '1' && mobile[1] >= '3' && mobile[1] <= '9' {
		return "880" + mobile, nil
	}

	// If it doesn't meet any of these criteria, return an error.
	return "", fmt.Errorf("Invalid mobile number format")
}
