package utils

import "strings"

func GetFileExtension(fileName string) string {
	split := strings.Split(fileName, ".")

	return split[len(split)-1]
}

func GetFileNameFromUrl(url string) string {
	split := strings.Split(url, "/")

	return split[len(split)-1]
}
