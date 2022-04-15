package main

import (
	"fmt"

	flatbuffers "github.com/google/flatbuffers/go"
)

func parseESM(buffer []byte) {
	buffer = buffer[78:]

	for len(buffer) > 0 {
		group := Group{}

		header := buffer[:24]

		group.Type = string(header[:4])
		group.Size = flatbuffers.GetUint32(header[4:8])
		group.Signature = string(header[8:12])

		fmt.Println("---------------")
		fmt.Println("GROUP: ", group.Signature)

		recordsBuffer := buffer[24:group.Size]

		for len(recordsBuffer) > 0 {
			record := Record{}

			header := recordsBuffer[:24]

			record.Type = string(header[:4])
			record.Size = flatbuffers.GetUint32(header[4:8])

			// fmt.Println("TYPE: ", record.Type)
			// fmt.Println("SIZE: ", record.Size)

			var offset uint32 = 24

			if record.Type == "GRUP" { // handling subgroups
				offset = 0

			}

			recordsBuffer = recordsBuffer[record.Size+uint32(offset):]
		}

		buffer = buffer[group.Size:]
	}
}
