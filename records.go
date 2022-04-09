package main

import flatbuffers "github.com/google/flatbuffers/go"

type record struct {
	s_type string
	size   uint32
	// contents  string
}

func newRecord(buffer []byte) record {
	header := buffer[:24]

	s_type := string(header[:4])
	size := flatbuffers.GetUint32(header[4:8])

	// contents := buffer[24:size]

	return record{
		s_type: s_type,
		size:   size,
		// contents:  string(contents),
	}
}
