package util

import (
	"fmt"
	"testing"
)

func TestPrint(t *testing.T) {
	// testCmd := "lpinfo"
	testCmd := "lp"
	out, _ := DoCmd(testCmd)
	
	fmt.Println(out.String())
}

func TestPrinterInfo(t *testing.T) {
	GetPrinter()
}
