package util

import (
	"bytes"
	"encoding/binary"
	"log"
)

func Handle(err error) {
	if err != nil {
		log.Panic(err)
	}
}

func ToHex(num int64) []byte {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, num)

	Handle(err)

	return buff.Bytes()
}
