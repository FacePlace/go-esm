package main

import (
	"encoding/binary"
	"fmt"
	"github.com/google/flatbuffers/go"
	"math"
)

func main() {
	buffer, err := readEsm("data/Skyrim/Skyrim.esm")

	if err != nil {
		panic(err)
	}

	fmt.Printf("%v\n", flatbuffers.GetUint8(buffer[0:1]))
	fmt.Printf("%v\n", flatbuffers.GetUint32(buffer[4:8]))
	fmt.Printf("%v", math.Float32frombits(binary.LittleEndian.Uint32(buffer[4:8])))
}
