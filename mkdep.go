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
	glen     = 16
	gTextLen = len("13A3F0F6-264A-3EF0-F2E0-DEC512342F34")
)

var (
	opcodes = map[string]byte{
		"TRUE": 6,
		"PUSH": 2,
		"AND":  3,
		"OR":   4,
		"END":  8,
	}
	// Intel's print format for GUIDs is stupid.
	// It could have been IPV6, but it's not.
	// They could have at least made
	// these fields fit in 64-bit ints,
	// they did not. Morons.
	// 13A3F0F6-264A-3EF0-F2E0-DEC512342F34
	// 000000000011111111112222222222333333
	// 012345678901234567890123456789012345
	fields = [...]struct {
		s, e int
	}{
		{0, 8},
		{9, 13},
		{14, 18},
		{19, 23},
		{24, 36},
	}
)

func toGuid(w io.Writer, s string) {
	//fmt.Printf("s %v len(s) %d\n", s, len(s))
	for l := len(s); l > 0; l = l - 2 {
		var b [1]byte
		//fmt.Printf("%v %d %d\n", s, l - 2, l + 0)
		if _, err := hex.Decode(b[:], []byte(s[l-2:l+0])); err != nil {
			log.Fatalf("err on %v: %v", s, err)
		}
		w.Write(b[:])
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
			if len(g) < gTextLen {
				log.Fatalf("'%v' is too short for a GUID, has to be %d chars\n", g, gTextLen)
			}
			// 13A3F0F6-264A-3EF0-F2E0-DEC512342F34
			for _, f := range fields {
				//fmt.Fprintf(os.Stderr, "%v %v\n", f.s, f.e)
				toGuid(&b, g[f.s:f.e])
			}
		}
	}
	l := b.Len() + 4
	hdr := append([]byte{byte(l), byte(l >> 8), byte(l >> 16), 0x13}, b.Bytes()...)
	if _, err := os.Stdout.Write(hdr); err != nil {
		log.Fatalf("%v", err)
	}

}
