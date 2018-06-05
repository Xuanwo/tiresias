package utils

import (
	"bufio"
	"io"
	"strings"

	"github.com/Xuanwo/tiresias/constants"
)

// Seek will seek to the start point of tiresias.
func Seek(r io.ReadSeeker) (cur int64, err error) {
	_, err = r.Seek(0, 0)
	if err != nil {
		return
	}

	cur = 0
	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		line := scanner.Text()

		// Ignore line not start with tiresias prefix or end with tiresias suffix.
		if !strings.HasPrefix(line, constants.CommentPrefix) || !strings.HasSuffix(line, constants.CommentSuffix) {
			// Current size should add len(len) and len("\n")
			cur += int64(len(scanner.Bytes())) + 1
			continue
		}

		return r.Seek(cur, 0)
	}

	return r.Seek(0, 2)
}
