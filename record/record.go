package record

import (
	"fmt"
	"os/exec"
	"sync"
)

// To work with json/txt files.
// Parses an array of data and starts writing it.
func Record(done chan struct{}, in Input, wg *sync.WaitGroup) {
	var commands []*exec.Cmd

	inputs := in.Parse()

	for i, input := range inputs {
		err := SingleRecord(input, &commands, i)
		if err != nil {
			continue
		}
	}

	go StopRecord(&commands, done, wg)
}

// Waits for a signal from the channel and then stops all recordings
func StopRecord(commands *[]*exec.Cmd, done chan struct{}, wg *sync.WaitGroup) {
	<-done
	defer wg.Done()

	for _, cmd := range *commands {
		if cmd.Process != nil {
			cmd.Process.Kill()
		}
	}

	fmt.Println("Recording of all streams has stopped")
}
