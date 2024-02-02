package parse

import (
	"bufio"
	"log"
	"os"
)

type FileTXT struct {
	TXTPath string
}

func (f FileTXT) Parse() []string {
	file, err := os.Open(f.TXTPath)
	if err != nil {
		log.Fatalf("can't open file %s", err)
	}
	defer file.Close()

	var urls []string
	read := bufio.NewScanner(file)

	for read.Scan() {
		urls = append(urls, read.Text())
	}

	if err := read.Err(); err != nil {
		log.Fatalf("wrong txt format %s", err)
	}

	return urls
}
