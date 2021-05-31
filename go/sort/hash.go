package sort

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"

//	"github.com/cespare/xxhash"
)
/*
func xxhash64(data Interface) uint64 {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.LittleEndian, data)
	if err != nil {
		log.Println("binary.Write failed:", err)
	}
	return xxhash.Sum64(buf.Bytes())
}
*/
func countHash(a interface{}) int {
	var hc int = 0
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.LittleEndian, a)
	if err != nil {
		log.Println("binary.Write failed:", err)
	}
	fmt.Printf("% x", buf.Bytes())
	data := buf.Bytes()
	for i := 0; i < len(data); i++ {
		hc += int(data[i] - '0')
	}
	return hc
}
