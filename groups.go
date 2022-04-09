package main

import flatbuffers "github.com/google/flatbuffers/go"

type group struct {
	s_type    string
	size      uint32
	signature string
	records   []record
}

func newGroup(buffer []byte) group {
	header := buffer[:24]

	s_type := string(header[:4])
	size := flatbuffers.GetUint32(header[4:8])
	signature := string(header[8:12])

	records_buffer := buffer[24:size]

	bytesConsumed := uint32(0)

	records := []record{}

	for bytesConsumed < uint32(len(records_buffer)) {
		r := newRecord(records_buffer[bytesConsumed:])

		bytesConsumed += r.size + 24

		records = append(records, r)
	}

	return group{
		s_type:    s_type,
		size:      size,
		signature: signature,
		records:   records,
		// contents:  string(contents),
	}
}
