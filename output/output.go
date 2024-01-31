package output

import (
	"fmt"
	"time"
)

func NameGenerate(i int) string {
	recordTime := time.Now().Format("20060102_150405")
	outputName := fmt.Sprintf("stream_%d_%s.mkv", i+1, recordTime)

	return outputName
}
