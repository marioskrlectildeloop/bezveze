package rd

import (
	"bufio"
	"io"
)

func Once(reader io.Reader) ([]byte, error) {
	r := bufio.NewReader(reader)

	b, _, err := r.ReadLine()
	if err != nil {
		return nil, err
	}

	return b, nil
}
