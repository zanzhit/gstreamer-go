package record

import (
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"

	"gst/output"
)

func SingleRecord(inputs string, commands *[]*exec.Cmd, i int) error {
	outputFile := output.NameGenerate(i)

	switch {
	case strings.Contains(inputs, " "):
		input := strings.Split(inputs, " ")
		videoFile := input[0]
		audioFile := input[1]

		parametres := fmt.Sprintf("gst-launch-1.0 rtspsrc location=%s ! rtph264depay ! h264parse ! matroskamux name=mux ! filesink location=%s rtspsrc location=%s ! rtpmp4gdepay ! aacparse ! mux.", videoFile, outputFile, audioFile)
		cmd, err := GstCommand(parametres, outputFile, i, commands)
		if err != nil {
			return err
		}

		go connectionTest(parametres, outputFile, i, commands, cmd)

		return nil

	case !strings.Contains(inputs, " "):
		parametres := fmt.Sprintf("gst-launch-1.0 rtspsrc location=%s ! rtph264depay ! h264parse ! matroskamux ! filesink %s", inputs, outputFile)
		cmd, err := GstCommand(parametres, outputFile, i, commands)
		if err != nil {
			return err
		}

		go connectionTest(parametres, outputFile, i, commands, cmd)

		return nil

	default:
		err := errors.New("wrong input type")

		return err
	}
}

func GstCommand(parametres string, outputFile string, i int, commands *[]*exec.Cmd) (*exec.Cmd, error) {
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

func connectionTest(parametres string, outputFile string, i int, commands *[]*exec.Cmd, cmd *exec.Cmd) {
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

		cmd, err := GstCommand(parametres, outputFile, i, commands)
		if err != nil {
			log.Printf("can't record %s\n", err)
			return
		}

		fmt.Printf("Recording the stream %d again\n", i+1)

		err = cmd.Start()
		if err != nil {
			log.Printf("Recording issue %s\n", err)
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
