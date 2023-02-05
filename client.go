package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

var (
	pipename = flag.String("pipename", "IPC_NAMED_PIPE", "Name of IPC pipe")
)

var ipc *IPC = &IPC{}

func main() {

	flag.Parse()

	if err := ipc.initWriter(); err != nil {
		log.Fatalf("Problem creating IPC named \"%s\": %v", ipc.fname, err)
	}

	log.Printf("Client ready to write to pipe named \"%s\" ...", ipc.fname)

	for {
		i := rand.Intn(32)
		if err := ipc.sendInt(i); err != nil {
			log.Printf("Unable to send number %d to monitor", i)
		}
		time.Sleep(time.Millisecond * 500)
	}
}

type IPC struct {
	file  *os.File
	fname string
}

func (ipc *IPC) initWriter() error {

	ipc.fname = *pipename

	var err error
	ipc.file, err = os.OpenFile(ipc.fname, os.O_WRONLY, os.ModeNamedPipe)
	return err
}

func (ipc *IPC) sendInt(i int) error {
	_, err := ipc.file.WriteString(fmt.Sprintf("%d\n", i))
	if err != nil {
		log.Printf("Unable to send data to monitor: %v", err)
		return err
	}

	return nil
}
