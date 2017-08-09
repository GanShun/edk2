package main

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"os"
)

const (
	glen = 16
)

var (
	opcodes = map[string]byte{
		"TRUE": 6,
		"PUSH": 2,
		"AND":  3,
		"OR":   4,
		"END":  8,
	}
	// 13A3F0F6-264A-3EF0-F2E0-DEC512342F34
	// 000000000011111111112222222222333333
	// 012345678901234567890123456789012345
	fields = [...]struct {
			gs, s, e int} {
			{0, 0, 8},
			{4, 9, 13},
			{6, 14, 18},
			{8, 19, 23},
			{10, 24, 36},
			}

)

func toGuid(w io.Writer, b []byte, s string) {
	for i := range b {
	}
}
func main() {
	var b bytes.Buffer

	for {
		var op, g string
		n, err := fmt.Scanln(&op, &g)
		if err == io.EOF {
			break
		}
		if n == 0 {
			continue
		}
		//fmt.Printf("%v %v\n", op, g)
		opcode, ok := opcodes[op]
		if !ok {
			log.Fatalf("Opcode %v not known", opcode)
		}
		b.Write([]byte{opcode})
		if op == "PUSH" {
			// 13A3F0F6-264A-3EF0-F2E0-DEC512342F34
			var d [glen]byte
			for i := range fields {
				if _, err := hex.Decode(d[fields[i].gs:], []byte(g[fields[i].s:fields[i].e])); err != nil {
					log.Fatalf("err on %v: %v", g[:8], err)
				}
			}

			b.Write(d[:])
		}
	}
	l := b.Len() + 4
	hdr := append([]byte{byte(l), byte(l>>8), byte(l>>16), 0x13}, b.Bytes()...)
	if _, err := os.Stdout.Write(hdr); err != nil {
		log.Fatalf("%v", err)
	}
	
}
