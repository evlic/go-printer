package cups

/*
#cgo LDFLAGS: -lcups
#include "cups/cups.h"
*/
import "C"

const printerStateIdle = "3"
const printerStatePrinting = "4"
const printerStateStopped = "5"

// PrinterState 打印机状态
type PrinterState string

const (
	idle     = PrinterState("3")
	printing = PrinterState("4")
	stopped  = PrinterState("5")
	
	// OSXTestFile >> resolve path/to/test/file
	// TODO: find the location of this file on linux/bsd/osx
	OSXTestFile = "/usr/share/cups/data/testprint"
)

var statusMapping = map[PrinterState]string{
	idle:     "闲置",
	printing: "工作中",
	stopped:  "停止工作",
}

// Printer 打印机数据结构
type Printer struct {
	Name      string
	IsDefault bool
	options   map[string]string // internal keys for status
	/*
		重要的属性
			job-sheets
	
			printer-state-reasons
			printer-is-accepting-jobs
			printer-make-and-model
	
			printer-state : 3
			printer-state-change-time : 1646993578
	*/
	/*
		打印机参数 >> "iCode_Printer"
			device-uri : dnssd://iCode-Printer%20%40%20icode.mac._ipps._tcp.local./?uuid=2421884b-10e5-3f9d-71b8-7bdaf798b0d2
			job-hold-until : no-hold
			job-sheets : none,none
			printer-info : iCode-Printer
			printer-state-reasons : none
			printer-type : 10653782
			copies : 1
			marker-change-time : 0
			number-up : 1
			printer-is-temporary : false
			printer-location : icode.mac
			printer-state-change-time : 1646993578
			job-cancel-after : 10800
			job-priority : 50
			printer-is-accepting-jobs : true
			printer-is-shared : false
			printer-make-and-model : HP LaserJet P2055 with Duplexer
			printer-state : 3
			finishings : 3
			printer-commands : AutoConfigure,Clean,PrintSelfTestPage
	*/
}

// Status returns the status of the dest
func (printer *Printer) Status() string {
	var returnMessage string
	
	// Return status of dest
	if value, ok := printer.options["printer-state"]; !ok {
		returnMessage = "系统未读取到打印机状态参数"
	} else if v, ok := statusMapping[PrinterState(value)]; ok {
		returnMessage = v
	}
	
	return returnMessage
}

// GetOptions returns a map of the dest options
func (printer *Printer) GetOptions() map[string]string {
	// Return option
	return printer.options
}

func (printer *Printer) TestPrint() int {
	var numOptions C.int
	var options *C.cups_option_t
	var jobID C.int
	
	file := OSXTestFile
	
	// Print a single file
	jobID = C.cupsPrintFile(
		C.CString(printer.Name), C.CString(file),
		C.CString("Test Print"), numOptions, options,
	)
	
	return int(jobID)
}

// Print 执行打印工作
func (printer *Printer) Print(file, jobName string) int {
	if CheckFilePath(file) {
		var (
			numOptions, jobID C.int
			options           *C.cups_option_t
		)
		
		jobID = C.cupsPrintFile(
			C.CString(printer.Name), C.CString(file),
			C.CString(jobName), numOptions, options,
		)
		
		return int(jobID)
	}
	return -1
}

// Jobs 检查当前打印机任务队列
func (printer *Printer) Jobs() {

}

// CheckFilePath TODO: 执行文件打印可行性校验
func CheckFilePath(file string) bool {
	return true
}
