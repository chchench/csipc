package main

import (
	"csipc"
	"log"
)

func main() {
	ipc := &csipc.CSIpc{}
	err := ipc.SetupIPCWriter("test")
	if err != nil {
		log.Printf("%v", err)
	}
}
