package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"math"
	"os"
	"os/signal"
	"syscall"

	"github.com/bsipos/thist"
)

var (
	pipename = flag.String("pipename", "IPC_NAMED_PIPE", "Name of IPC pipe")
)

func dataStream() chan float64 {
	c := make(chan float64, 1024)
	go runStreamReader(c)
	return c
}

var total int

func main() {

	flag.Parse()

	installSigIntHandler()

	log.Printf("Monitor starting up ...")

	h := thist.NewHist(nil, "Retrieval Size Histogram - Log2(Size) = X", "fixed", 32, false)
	c := dataStream()

	log.Printf("Monitor ready to read from pipe named \"%s\" ...", *pipename)

	total := 0
	for {
		v := <-c

		h.Update(v)
		if total%50 == 0 {
			fmt.Println(h.Draw())
		}

		total++
	}
}

func installSigIntHandler() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		cleanUp()
		os.Exit(1)
	}()
}

func cleanUp() {
	log.Printf("A total of %d numbers were received\n", total)
	os.Remove(*pipename)
	os.Exit(1)
}

func runStreamReader(c chan float64) error {

	name := *pipename

	os.Remove(name)
	err := syscall.Mkfifo(name, 0666)
	if err != nil {
		return err
	}

	f, err := os.OpenFile(name, os.O_RDONLY, os.ModeNamedPipe)
	if err != nil {
		return err
	}

	reader := bufio.NewReader(f)

	var value int

	for {
		line, err := reader.ReadBytes('\n')
		if err == nil {

			fmt.Sscanf(string(line), "%d\n", &value)
			c <- math.Ceil(float64(value))
		}
	}

	return nil
}
