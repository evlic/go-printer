package util

import (
	"bytes"
	"os/exec"
)

func DoCmd(name string, args ...string) (out *bytes.Buffer, ok bool) {
	out = &bytes.Buffer{}
	cmd := exec.Command(name, args...)
	cmd.Stdout, cmd.Stderr = out, out
	cmd.Start()

	// fmt.Printf("cmd.ProcessState.SysUsage(): %v\n", cmd.ProcessState.SysUsage())

	if e := cmd.Wait(); e != nil {
		ok = false
	} else {
		ok = true
	}

	return
}

// func RunCommand(name string, arg ...string) (ok bool, out *io.ReadCloser, q chan byte) {
// 	cmd := exec.Command(name, arg...)
// 	stdout, err := cmd.StdoutPipe()
// 	if err != nil {
// 		ok = false
// 	}

// 	// 重定向错误输出 到 标准输出
// 	cmd.Stderr = cmd.Stdout

// 	out = &stdout

// 	q = make(chan byte)

// 	go func() {
// 		cmd.Wait()
// 		q <- 'q'
// 	}()

// }
