package util

import (
	"fmt"
	"os"
	"strings"
	"testing"
	"time"
)

func TestLogFormat(t *testing.T) {
	for i := 0; i < 100; i++ {
		time.Sleep(50 * time.Millisecond)
		word := strings.Repeat("=", i) + strings.Repeat(" ", 99-i)
		fmt.Printf("\r%.0f%%[%s]", float64(i)/99*100, color(word, i))
		os.Stdout.Sync()
	}
}

func color(word string, i int) string {
	return fmt.Sprintf("%c[%d;%d;%dm%s%c[0m", 0x1B, 1, 0, 30+(i%7), word, 0x1B)
}
