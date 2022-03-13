package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {

	cnt := 100

	for i := 0; i < cnt; i++ {
		time.Sleep(100 * time.Millisecond)

		word := strings.Repeat("=", i) + strings.Repeat(" ", cnt-1-i)

		fmt.Printf("\r%.0f%%[%s]", float64(i)/99*100, color(word, i))
		_ = os.Stdout.Sync()
	}
}

func color(word string, i int) string {
	return fmt.Sprintf("%c[%d;%d;%dm%s%c[0m", 0x1B, 1, 0, 31+(i%7), word, 0x1B)
}
