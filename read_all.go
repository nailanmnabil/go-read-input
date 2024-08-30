package goreadinput

import (
	"io"
	"log"
	"os"
)

func ReadAll(name string) {
	file, err := os.Open(name)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	_, err = io.ReadAll(file)
	if err != nil {
		log.Fatalln(err)
	}
}
