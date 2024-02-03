package record

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"

	"gst/output"
)

// Starts recording the line that was read and selects the required recording parameters.
func SingleRecord(inputs string, commands *[]*exec.Cmd, i int) error {
	outputFile := output.NameGenerate(i)

	switch {
	case strings.Contains(inputs, " "):
		cmd, parametres, err := TwoInputsArg(inputs, outputFile, i, commands)
		if err != nil {
			return err
		}

		go ConnectionTest(parametres, outputFile, i, commands, cmd)

		return nil

	default:
		cmd, parametres, err := SingleInputArg(inputs, outputFile, i, commands)
		if err != nil {
			return err
		}

		go ConnectionTest(parametres, outputFile, i, commands, cmd)

		return nil
	}
}

// Generates and runs the gstreamer console command
func gstCommand(parametres string, outputFile string, i int, commands *[]*exec.Cmd) (*exec.Cmd, error) {
	c := strings.Split(parametres, " ")
	cmd := exec.Command(c[0], c[1:]...)
	*commands = append(*commands, cmd)

	fmt.Printf("Record of stream %d has started.\n", i+1)

	err := cmd.Start()
	if err != nil {
		log.Printf("can't start recording %s\n", err)
	}

	return cmd, err
}

// Checks if the file is being written using the file weight
func ConnectionTest(parametres string, outputFile string, i int, commands *[]*exec.Cmd, cmd *exec.Cmd) {
	time.Sleep(time.Second * 4)

	fileInfo, err := os.Stat(outputFile)
	if err != nil {
		log.Printf("stream check availability issue %s\n", err)
		return
	}

	for fileInfo.Size() == 0 {
		if cmd.Process != nil {
			cmd.Process.Kill()
		}

		_, err := gstCommand(parametres, outputFile, i, commands)
		if err != nil {
			log.Printf("can't record %s\n", err)
			return
		}

		time.Sleep(time.Second * 4)

		fileInfo, err = os.Stat(outputFile)
		if err != nil {
			log.Printf("stream check availability issue, %s\n", err)
			return
		}
	}
}
