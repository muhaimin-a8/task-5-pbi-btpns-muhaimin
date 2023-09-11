package main

import (
	"strings"
	"testing"
)

func TestName(t *testing.T) {
	name := "1_Ifpd_HtDiK9u6.h68SZ.dsds.dsdsd.dsdsd.dampodvmosdv.sdidnioas.gNuA.JPG"

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
