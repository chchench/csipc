package main

import (
	"os"
)

func main() {
	data := []float64{
		0.1,
		0.2, 0.21, 0.22, 0.22,
		0.3,
		0.4,
		0.5, 0.51, 0.52, 0.53, 0.54, 0.55, 0.56, 0.57, 0.58,
		0.6,
		// 0.7 is empty
		0.8,
		0.9,
		1.0,
	}

	bins := 9
	hist := Hist(bins, data)

	maxWidth := 5
	err := os.Fprint(os.Stdout, hist, Linear(maxWidth))
}
	