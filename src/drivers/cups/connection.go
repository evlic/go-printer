package cups

/*
#cgo LDFLAGS: -lcups
#include "cups/cups.h"
*/
import "C"
import (
	"unsafe"
)

// Connection is a struct containing information about a CUPS connection
type Connection struct {
	isDefault bool
	address   string
	NumDests  int
	Dests     []*Printer
}

// Is there a better way to do this?
// I manually calculated the size in C
//
const cupsDestTSize = 32
const cupsOptionTSize = 16

var enumDestsChan chan *Printer

// Refresh updates the list of destinations and their state
func (c *Connection) Refresh() {
	updateDefaultConnection(c)
}

// NewDefaultConnection 创建新的 CUPS 连接
func NewDefaultConnection() *Connection {

	connection := &Connection{
		isDefault: true,
	}
	updateDefaultConnection(connection)

	return connection
}

// updateDefaultConnection 更新默认的 CUPS 连接对象, 核心方法，调用 CUPS 函数
// TODO: 检查内存泄露
func updateDefaultConnection(c *Connection) {
	var dests *C.cups_dest_t
	destCount := C.cupsGetDests(&dests)
	goDestCount := int(destCount)

	var destsArr []*Printer

	destPtr := uintptr(unsafe.Pointer(dests))
	for i := 0; i < goDestCount; i++ {

		// is this ok?
		dest := (*C.cups_dest_t)(unsafe.Pointer(destPtr))
		d := &Printer{
			Name: C.GoString(dest.name),
		}

		destOptPtr := uintptr(unsafe.Pointer(dest.options))

		options := make(map[string]string)
		for j := 0; j < int(dest.num_options)-1; j++ {
			opt := (*C.cups_option_t)(unsafe.Pointer(destOptPtr))
			options[C.GoString(opt.name)] = C.GoString(opt.value)
			destOptPtr += uintptr(cupsOptionTSize)
		}

		d.options = options

		destsArr = append(destsArr, d)
		destPtr += uintptr(cupsDestTSize)
	}

	// free the pointers
	C.cupsFreeDests(destCount, dests)

	c.NumDests = goDestCount
	c.Dests = destsArr
}
