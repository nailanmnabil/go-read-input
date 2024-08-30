package goreadinput

import (
	"bufio"
	"io"
	"log"
	"os"
)

func BufferedReading(name string) {
	file, err := os.Open(name)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	
	for {
		buffer := make([]byte, 1024)
		_, err := reader.Read(buffer)
		
		// _, err := file.Read(buffer) // alternative

		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatalln(err)
		}
	}
}
