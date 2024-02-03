package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"sync"

	"gst/parse"
	"gst/record"
)

func main() {
	wg := new(sync.WaitGroup)
	done := make(chan struct{})

	fmt.Print("Enter input data:")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()

	go func() {
		var in string
		fmt.Scanln(&in)
		close(done)
	}()

	wg.Add(1)
	inputCheck(input, done, wg)
	wg.Wait()

	fmt.Println("Program completed.")
}

//НАЗВАНИЯ ФАЙЛОВ НЕ ДОЛЖНЫ СОДЕРЖАТЬ ПРОБЕЛОВ

func inputCheck(input string, done chan struct{}, wg *sync.WaitGroup) {
	switch {
	case strings.Contains(input, ".txt"):
		file := parse.FileTXT{TXTPath: input}
		record.Record(done, file, wg)

	case strings.Contains(input, ".json"):
		file := parse.FileJSON{JSONPath: input}
		record.Record(done, file, wg)

	default:
		var commands []*exec.Cmd

		err := record.SingleRecord(input, &commands, 0)
		if err != nil {
			log.Fatalf("can't start recording %s", err)
		}

		go record.StopRecord(&commands, done, wg)
	}
}
