package csipc

import (
	"bufio"
	"fmt"
	"os"
	"syscall"
	"time"
)

type CSIpc struct {
	fn string
}

func (ipc *CSIpc) SetupIPCReader(name string) error {

	ipc.fn = genPipeFilename(name)
	os.Remove(ipc.fn)

	err := syscall.Mkfifo(ipc.fn, 0666)
	if err != nil {
		return err
	}

	f, err := os.OpenFile(ipc.fn, os.O_RDONLY, os.ModeNamedPipe)
	if err != nil {
		return err
	}

	reader := bufio.NewReader(f)

	for {
		line, err := reader.ReadBytes('\n')
		if err == nil {
			fmt.Print("SERVER: reading data:  " + string(line))
		}
	}
}

func (ipc *CSIpc) SetupIPCWriter(name string) error {

	ipc.fn = genPipeFilename(name)

	f, err := os.OpenFile(ipc.fn, os.O_WRONLY, os.ModeNamedPipe)
	if err != nil {
		return err
	}

	i := 0

	for {
		f.WriteString(fmt.Sprintf("CLIENT: sending count %d\n", i++))
		time.Sleep(time.Second * 5)
	}
}

func genPipeFilename(prefix string) string {
	return prefix + prefix
}
