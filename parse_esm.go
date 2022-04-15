package main

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

func parseESM(buffer []byte) {
	buffer = buffer[78:]

	groups := []Group{}

	for len(buffer) > 0 {
		group := Group{}

		header := buffer[:24]

		group.Type = string(header[:4])
		group.Size = flatbuffers.GetUint32(header[4:8])
		group.Signature = string(header[8:12])
		group.Records = []Record{}

		// fmt.Println("---------------")
		// fmt.Println("GROUP: ", group.Signature)

		recordsBuffer := buffer[24:group.Size]

		for len(recordsBuffer) > 0 {
			record := Record{}

			header := recordsBuffer[:24]

			record.Type = string(header[:4])
			record.Size = flatbuffers.GetUint32(header[4:8])
			record.Subrecords = []Subrecord{}

			var offset uint32 = 24

			if record.Type == "GRUP" { // handling subgroups
				offset = 0
			} else if record.Type != "NPC_" && record.Type != "WRLD" {
				subrecordsBuffer := recordsBuffer[offset : record.Size+offset]

				for len(subrecordsBuffer) > 0 {
					subrecord := Subrecord{}
					header := subrecordsBuffer[:6]

					subrecord.Type = string(header[:4])
					subrecord.Size = flatbuffers.GetUint16(header[4:6])

					subrecord.Buffer = string(subrecordsBuffer[6 : 6+subrecord.Size])

					subrecordsBuffer = subrecordsBuffer[6+subrecord.Size:]

					record.Subrecords = append(record.Subrecords, subrecord)
				}
			}

			recordsBuffer = recordsBuffer[record.Size+uint32(offset):]

			group.Records = append(group.Records, record)
		}

		buffer = buffer[group.Size:]

		groups = append(groups, group)
		// if group.Signature == "WEAP" {
		// 	fmt.Printf("%v\n", group)
		// }
	}
}
