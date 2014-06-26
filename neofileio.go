package neo

import (
	"bufio"
	"io"
	"os"
)

func ReadFile(fileName string, lineReceiver func(line string) error) error {
	f, err := os.Open(os.Args[1])
	if err != nil {
		return err
	}
	defer f.Close()

	r := bufio.NewReader(f)
	for {
		line, err := r.ReadString('\n')

		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}

		err = lineReceiver(line)
		if err != nil {
			return err
		}
	}
}

type NullReaderWriterCloser struct {
}

func (n NullReaderWriterCloser) Read(p []byte) (int, error) {
	return 0, io.EOF
}

func (n NullReaderWriterCloser) Write(p []byte) (int, error) {
	return len(p), nil
}

func (n NullReaderWriterCloser) Close() error {
	return nil
}
