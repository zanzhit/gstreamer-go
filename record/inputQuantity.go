package record

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
)

// Start recording in the presence of two arguments. Checks and selects rtsp/mkv.
func TwoInputsArg(inputs string, outputFile string, i int, commands *[]*exec.Cmd) (*exec.Cmd, string, error) {
	var parametres string

	input := strings.Split(inputs, " ")
	videoFile := input[0]
	audioFile := input[1]

	if strings.Contains(videoFile, ".mkv") {
		parametres = fmt.Sprintf("gst-launch-1.0 filesrc location=%s ! matroskademux name=demux filesrc location=%s ! matroskademux name=demux2 demux.video_0 ! queue ! mux.video_0 demux2.audio_0 ! queue ! mux.audio_0 matroskamux name=mux ! filesink location=out%d.mkv", videoFile, audioFile, i)
	} else {
		parametres = fmt.Sprintf("gst-launch-1.0 rtspsrc location=%s ! rtph264depay ! h264parse ! matroskamux name=mux ! filesink location=%s rtspsrc location=%s ! rtpmp4gdepay ! aacparse ! mux.", videoFile, outputFile, audioFile)
	}

	cmd, err := gstCommand(parametres, outputFile, i, commands)
	if err != nil {
		log.Printf("gstreamer wrong data %s", err)
		return nil, "", err
	}

	return cmd, parametres, nil
}

// Start recording for one rtsp camera
func SingleInputArg(inputs string, outputFile string, i int, commands *[]*exec.Cmd) (*exec.Cmd, string, error) {
	parametres := fmt.Sprintf("gst-launch-1.0 rtspsrc location=%s ! rtph264depay ! h264parse ! matroskamux ! filesink location=%s", inputs, outputFile)

	cmd, err := gstCommand(parametres, outputFile, i, commands)
	if err != nil {
		log.Printf("gstreamer wrong data %s", err)
		return nil, "", err
	}

	return cmd, parametres, nil
}
