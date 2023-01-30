package main

import (
	"csipc"
	"log"
)

func main() {
	ipc := &csipc.CSIpc{}
	err := ipc.SetupIPCReader("test")
	if err != nil {
		log.Printf("%v", err)
	}
}
