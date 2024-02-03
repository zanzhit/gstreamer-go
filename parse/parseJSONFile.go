package parse

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type FileJSON struct {
	JSONPath string
}

// Opens file.
func readFile(fileName string) ([]byte, error) {
	r, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	if !json.Valid(r) {
		return nil, fmt.Errorf("invalid json")
	}

	return r, nil
}

// Decodes the required json fields.
func decodeJSON(f []byte) ([]string, error) {
	var config Config
	err := json.Unmarshal(f, &config)
	if err != nil {
		return nil, err
	}

	var data []string

	for _, j := range config.ArraySources {
		d := fmt.Sprintf("%s %s", j.VideoSource, j.AudioSource)
		data = append(data, d)
	}
	return data, nil
}

// Combines file opening with json parsing.
func (f FileJSON) Parse() []string {
	file, err := readFile(f.JSONPath)
	if err != nil {
		log.Fatalf("can't open file %s", err)
	}

	data, err := decodeJSON(file)
	if err != nil {
		log.Fatalf("wrong json %s", err)
	}

	return data
}
