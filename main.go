package main

import (
	"bufio"
	"fmt"
	"os"
)

func readEsm(filename string) ([]byte, error) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	stats, statsErr := file.Stat()
	if statsErr != nil {
		return nil, statsErr
	}

	var size int64 = stats.Size()
	bytes := make([]byte, size)

	buffer := bufio.NewReader(file)
	_, err = buffer.Read(bytes)
	if err != nil {
		panic(err)
	}

	return bytes, err
}

func main() {
	buffer, err := readEsm("data/Skyrim/Skyrim.esm")

	if err != nil {
		panic(err)
	}

	// fmt.Printf("%v\n", binary.LittleEndian.Uint32(buffer[4:8]))
	// fmt.Printf("%v", math.Float32frombits(binary.LittleEndian.Uint32(buffer[4:8])))

	fmt.Printf("%s", buffer[:24])
}
