package tail

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func Tail(file string, lines int) (*os.File, error) {
	r, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	if lines <= 0 {
		return r, nil
	}
	if err := tail(r, lines); err != nil {
		return nil, err
	}
	return r, nil
}

var SeekStep int64 = 1 << 12

func tail(r io.ReadSeeker, lines int) error {
	var (
		buf    []byte
		step   int64
		seek   int64
		offset int64
		err    error
	)
	if seek, err = r.Seek(0, io.SeekEnd); err != nil || lines == 0 {
		return err
	}
	offset = seek

	if seek >= SeekStep {
		buf, step = make([]byte, SeekStep), SeekStep
	} else {
		buf, step = make([]byte, seek), seek
	}
	lines++
	for lines > 0 {
		if seek, err = seekReader(r, seek-step); err != nil {
			return fmt.Errorf("seek: %s", err)
		}
		if _, err := io.ReadFull(r, buf); err != nil {
			return fmt.Errorf("read: %s", err)
		}
		lines -= bytes.Count(buf, []byte{'\n'})
		if seek == 0 {
			break
		}
	}

	if k, _ := r.Seek(0, io.SeekCurrent); k == offset {
		if seek, err = seekReader(r, seek-step); err != nil {
			return fmt.Errorf("reset: %s", err)
		}
	}
	lines--
	if lines < 0 {
		seek, err = r.Seek(seek, io.SeekStart)
		if err != nil {
			return err
		}
	}
	for lines < 0 {
		x := bytes.IndexByte(buf, '\n')
		if x < 0 {
			break
		}
		seek += int64(x) + 1
		buf = buf[x+1:]
		lines++
	}
	if seek, err = r.Seek(seek, io.SeekStart); err != nil {
		return err
	}
	return err
}

func seekReader(r io.Seeker, offset int64) (int64, error) {
	if offset < 0 {
		offset = 0
	}
	return r.Seek(offset, io.SeekStart)
}
