package tail

import (
	"bufio"
	"io"
	"strings"
	"testing"
)

var sample = []string{
	"line #01",
	"line #02",
	"line #03",
	"line #04",
	"line #05",
	"line #06",
	"line #07",
	"line #08",
	"line #09",
	"line #10",
	"line #11",
	"line #12",
	"line #13",
	"line #14",
	"line #15",
}

func TestTail(t *testing.T) {
	var (
		text = strings.Join(sample, "\n")
		size = len(sample) - 1
	)
	for i := 1; i < len(sample); i++ {
		r := strings.NewReader(text)
		if err := tail(r, i); err != nil {
			t.Errorf("error while tail-ing sample: %s", err)
			continue
		}
		compareLines(t, r, sample[size-i:])
	}
}

func compareLines(t *testing.T, r io.Reader, lines []string) {
	t.Helper()
	scan := bufio.NewScanner(r)
	for i := range lines {
		if !scan.Scan() {
			t.Errorf("reader is empty! want %s", lines[i])
			return
		}
		if lines[i] != scan.Text() {
			t.Errorf("lines mismatched! want: %q, got %q", lines[i], scan.Text())
			break
		}
	}
	for scan.Scan() {
		t.Errorf("remaining line: %s", scan.Text())
	}
}
