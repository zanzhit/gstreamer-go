package record

import (
	"fmt"
	"os/exec"
)

func Record(done chan struct{}, in Input) {
	var commands []*exec.Cmd

	inputs := in.Parse()

	for i, input := range inputs {
		err := SingleRecord(input, &commands, i)
		if err != nil {
			continue
		}
	}

	go stopRecord(&commands, done)
}

func stopRecord(commands *[]*exec.Cmd, done chan struct{}) {
	<-done

	for i, cmd := range *commands {
		if cmd.Process != nil {
			cmd.Process.Kill()
			fmt.Printf("Stream %d recording interrupted", i+1)
		}
	}
}
