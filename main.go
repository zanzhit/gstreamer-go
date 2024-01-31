package main

import (
	"fmt"

	"gst/record"
	"gst/record/gsttxt"
)

func main() {
	file := gsttxt.FileTXT{TXTPath: "urls.txt"}

	done := make(chan struct{})

	go func() {
		var input string
		fmt.Scanln(&input)
		close(done)
	}()

	record.Record(done, file)

	<-done

	fmt.Println("Программа завершена")
}
