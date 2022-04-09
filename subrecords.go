package main

import flatbuffers "github.com/google/flatbuffers/go"

type subrecord struct {
	s_type   string
	size     uint16
	contents string
}

func newSubrecord(buffer []byte) subrecord {
	header := buffer[:6]

	s_type := string(header[:4])
	size := flatbuffers.GetUint16(header[4:6])

	contents := buffer[6 : 6+size]

	return subrecord{
		s_type:   s_type,
		size:     size,
		contents: string(contents),
	}
}
