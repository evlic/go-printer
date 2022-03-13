package cups

import (
	"fmt"
	"testing"
)

func PrintMap(m map[string]string) {
	for k, v := range m {
		fmt.Printf("\t%v : %v\n", k, v)
	}
}

func TestPrinter_GetOverView(t *testing.T) {
	// 获得默认打印机连接对象
	printers := NewDefaultConnection()
	for _, printer := range printers.Dests {
		fmt.Printf("%v (%v)\n", printer.Name, printer.Status())
	}
}

func TestPrinter_GetOptions(t *testing.T) {
	// 获得默认打印机连接对象
	printers := NewDefaultConnection()
	for _, printer := range printers.Dests {
		fmt.Printf("打印机参数 >> %#v\n", printer.Name)
		PrintMap(printer.GetOptions())
	}
}

func TestPrint(t *testing.T) {
	// file := "xxx"
	printers := NewDefaultConnection()
	printer := printers.Dests[0]
	PrintMap(printer.options)
}
