package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"
)

var (
	pipename = flag.String("pipename", "IPC_NAMED_PIPE", "Name of IPC pipe")
	monitor  = flag.Bool("monitor", false, "True to send data to monitor")
)

var ipc *IPC = &IPC{}

func main() {

	flag.Parse()

	if err := ipc.initWriter(); err != nil {
		log.Fatalf("Problem creating IPC named \"%s\": %v", ipc.fname, err)
	}

	log.Printf("Client ready to write to pipe named \"%s\" ...", ipc.fname)

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

func gatherFilePaths(p string) []string {

	var fileList []string

	fileInfo, err := os.Stat(p)
	if err != nil {
		log.Fatalf("The path %s does not appear to be a valid one.", p)
	}

	if fileInfo.IsDir() {
		files, err := ioutil.ReadDir(p)
		if err != nil {
			log.Fatal(err)
		}

		for _, f := range files {
			if strings.HasPrefix(f.Name(), ".") {
				continue
			}

			filePath := path.Join(p, f.Name())

			fileInfo, err := os.Stat(filePath)
			if err != nil || fileInfo.IsDir() {
				continue
			}

			fileList = append(fileList, filePath)
		}
	} else {
		fileList = append(fileList, p)
	}

	return fileList
}
