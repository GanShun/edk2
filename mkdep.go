package main

import (
	"bytes"
	"fmt"
	"encoding/hex"
	"io"
	"log"
	"os"
)

const (
	glen = 16
)

var (
	opcodes = map[string]byte {
 "TRUE": 6,
 "PUSH": 2,
 "AND": 3,
 "OR": 4,
 "END": 8,
	}
)


func tobin(r io.Write, []byte) {
	var b [16]byte
	if n, err := r.Read(b[:]); n != 16 || err != nil {
		log.Fatalf("Reading GUID: got %d of 16 bytes: %v", n, err)
	}
	return b[:]
}

func main() {
	var b bytes.Buffer
	var (
		hdr [4]byte
		len int
	)

	for {
		var op, val string
		if n, err := fmt.Fscanf(os.Stdin, &op); err != nil || n == 0 {
			log.Fatalf("Scanln: Scanlf: no opcode: %v", err)
		}

		opcode, ok := opcodes[op]
		if ! ok {
			log.Fatalf("Opcode %v not known", opcode)
		}
		if opcode == "PUSH" {
			if n, err := fmt.Fscanf(os.Stdin, &guid); err != nil || n == 0 {
				log.Fatalf("Scanln: Scanlf: no guid: %v", err)
			}
		}
		// 13A3F0F6-264A-3EF0-F2E0-DEC512342F34
		if _, err := hex.Decode(d[0], g[:8]); err != nil {
			log.Fatalf("err on %v: %v", g[:8], err)
		}
		
		fmt.Printf("write %v %v\n", op, guid)
	}
}
