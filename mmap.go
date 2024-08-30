package goreadinput

import (
	"log"
	"os"
	"syscall"
)

func Mmap(name string) {
	file, err := os.Open(name)
	if err != nil {
		log.Fatalln(err)
	}

	stat, err := file.Stat()
	if err != nil {
		log.Fatalln(err)
	}

	data, err := syscall.Mmap(int(file.Fd()), 0, int(stat.Size()), syscall.PROT_READ, syscall.MAP_SHARED)
	if err != nil {
		log.Fatalln(err)
	}
	defer func() {
		err = syscall.Munmap(data)
		if err != nil {
			log.Fatalln(err)
		}
	}()
}
