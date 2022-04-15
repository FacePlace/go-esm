package main

type HeaderSize int

type Group struct {
	Type      string
	Size      uint32
	Signature string
	Records   []Record
}

type Record struct {
	Type       string
	Size       uint32
	Subrecords []Subrecord
}

type Subrecord struct {
	Type   string
	Size   uint16
	Buffer string
}
