package util

import (
	"io/ioutil"
	"log"
)

func ReadInput(path string) string {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalf("could not open %s: %v", path, err)
	}
	return string(data)
}
