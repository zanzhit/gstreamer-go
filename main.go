package main

import (
	"fmt"
	"sync"

	"gst/record"
	"gst/record/gsttxt"
)

func main() {
	wg := new(sync.WaitGroup)

	file := gsttxt.FileTXT{TXTPath: "urls.txt"}

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
