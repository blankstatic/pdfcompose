package utils

import "os"

func GetVar(envVar, defaultVal string) string {
	value := os.Getenv(envVar)
	if value == "" {
		value = defaultVal
	}
	return value

}
