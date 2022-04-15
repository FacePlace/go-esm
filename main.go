package main

import (
	"fmt"
	"time"
)

func main() {
	performance := time.Now()

	buffer, err := readEsm("data/Skyrim/Skyrim.esm")

	if err != nil {
		panic(err)
	}

	parseESM(buffer)

	fmt.Printf("Finished processing in %v\n", time.Since(performance))
}
