package record

import (
	"fmt"
	"os/exec"
	"sync"
)

func Record(done chan struct{}, in Input, wg *sync.WaitGroup) {
	var commands []*exec.Cmd

	inputs := in.Parse()

	for i, input := range inputs {
		err := SingleRecord(input, &commands, i)
		if err != nil {
			continue
		}
	}

	go stopRecord(&commands, done, wg)
}

func stopRecord(commands *[]*exec.Cmd, done chan struct{}, wg *sync.WaitGroup) {
	<-done
	defer wg.Done()

	for _, cmd := range *commands {
		if cmd.Process != nil {
			cmd.Process.Kill()
		}
	}

	fmt.Println("Recording of all streams has stopped")
}
