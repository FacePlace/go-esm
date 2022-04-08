package main

import (
	"encoding/binary"
	"fmt"
	"math"
)

func main() {
	buffer, err := readEsm("data/Skyrim/Skyrim.esm")

	if err != nil {
		panic(err)
	}

	fmt.Printf("%v\n", binary.LittleEndian.Uint32(buffer[4:8]))
	fmt.Printf("%v", math.Float32frombits(binary.LittleEndian.Uint32(buffer[4:8])))
}
