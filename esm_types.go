package main

type HeaderSize int

type Group struct {
	Type      string
	Size      uint32
	Signature string
	records   []Record
}

type Record struct {
	Type       string
	Size       uint32
	subrecords []Subrecord
}

type Subrecord struct {
	Type   string
	Size   uint16
	buffer []byte
}
