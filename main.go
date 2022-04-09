package main

import (
	"fmt"
)

func main() {
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

		fmt.Printf("%v\n", g)
	}
	fmt.Printf("%v\n", len(groups))
}
