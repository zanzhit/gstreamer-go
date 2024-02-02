package main

import (
	"fmt"
	"sync"

	"gst/parse"
	"gst/record"
)

func main() {
	wg := new(sync.WaitGroup)

	file := parse.FileJSON{JSONPath: "urls.json"}

	done := make(chan struct{})

	go func() {
		var input string
		fmt.Scanln(&input)
		close(done)
	}()

	wg.Add(1)
	record.Record(done, file, wg)

	wg.Wait()

	fmt.Println("Программа завершена")
}
