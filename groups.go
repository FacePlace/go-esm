package main

import flatbuffers "github.com/google/flatbuffers/go"

type group struct {
	s_type    string
	size      uint32
	signature string
	// contents  string
}

func newGroup(buffer []byte) group {
	header := buffer[:24]

	s_type := string(header[:4])
	size := flatbuffers.GetUint32(header[4:8])
	signature := string(header[8:12])

	// contents := buffer[24:size]

	return group{
		s_type:    s_type,
		size:      size,
		signature: signature,
		// contents:  string(contents),
	}
}
