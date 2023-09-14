package main

import (
	"fmt"
	"pbi-btpns-api/utils"
	"strings"
	"testing"
)

func TestName(t *testing.T) {
	//name := "1_Ifpd_HtDiK9u6.h68SZ.dsds.dsdsd.dsdsd.dampodvmosdv.sdidnioas.gNuA.JPG"
	url := "http://localhost:8080/static/photos/075bf988-681b-4df1-b3b7-222eb0ba4d45.png"
	fmt.Println(utils.GetFileNameFromUrl(url))

}

func isFilePermitted(fileName string) bool {
	prefix := []string{"png", "jpg", "jpeg"}
	for _, v := range prefix {
		if strings.HasSuffix(strings.ToLower(fileName), "."+v) {
			return true
		}
	}

	return false
}
