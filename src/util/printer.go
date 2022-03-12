package util

import "fmt"

func GetPrinter() {
	if out, ok := DoCmd("lpinfo", "-v"); ok {
		fmt.Println(out.String())
	}
	
}
