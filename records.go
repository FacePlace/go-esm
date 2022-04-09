package main

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type record struct {
	s_type     string
	size       uint32
	subrecords []subrecord
}

func newRecord(buffer []byte) record {
	header := buffer[:24]

	s_type := string(header[:4])
	size := flatbuffers.GetUint32(header[4:8])

	subrecords_buffer := buffer[24 : 24+size]

	bytesConsumed := uint16(0)

	subrecords := []subrecord{}

	// fmt.Printf("%v\n", s_type)
	// fmt.Printf("-\n")
	for bytesConsumed < uint16(len(subrecords_buffer)) {
		s := newSubrecord(subrecords_buffer[bytesConsumed:])
		// fmt.Println(s)

		bytesConsumed += s.size + 6

		subrecords = append(subrecords, s)
	}
	// fmt.Printf("---\n")

	return record{
		s_type:     s_type,
		size:       size,
		subrecords: subrecords,
	}
}
