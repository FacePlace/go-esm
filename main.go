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

	var bytesConsumed uint32 = 78
	groups := []group{}

	for bytesConsumed < uint32(len(buffer)) {
		g := newGroup(buffer[bytesConsumed:])

		bytesConsumed += g.size

		groups = append(groups, g)

		if g.signature == "WEAP" {
			// fmt.Printf("%v\n", g.records)
		}
	}
	// fmt.Printf("%v\n", len(groups))

	fmt.Printf("Finished processing in %v\n", time.Since(performance))
}
