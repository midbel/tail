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
	data := []struct {
		Lines int
		Want  []string
	}{
		{
			Lines: 1,
			Want:  sample[14:],
		},
		{
			Lines: 20,
			Want:  sample[0:],
		},
		{
			Lines: 10,
			Want:  sample[4:],
		},
	}
	text := strings.Join(sample, "\n")
	for _, d := range data {
		t.Logf("tail: %d lines", d.Lines)
		r := strings.NewReader(text)
		if err := tail(r, d.Lines); err != nil {
			t.Errorf("error while tail-ing sample: %s", err)
			continue
		}
		compareLines(t, r, d.Want)
	}
}

func compareLines(t *testing.T, r io.Reader, lines []string) {
	t.Helper()
	scan := bufio.NewScanner(r)
	for i := 0; scan.Scan(); i++ {
		if i >= len(lines) {
			t.Errorf("too many lines generated! want %d, got %d", len(lines), i)
			break
		}
		if lines[i] != scan.Text() {
			t.Errorf("lines mismatched! want: %q, got %q", lines[i], scan.Text())
			break
		}
	}
}
