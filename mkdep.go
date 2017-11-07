package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
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
	stackdepth int
)

func writeGUID(w io.Writer, guid string) {
	s := strings.Split(guid, "-")
	//fmt.Fprintf(os.Stderr, "%v %v: ", guid, s)
	bits := []int{32, 16, 16, 16, 8, 8, 8, 8, 8, 8}

	// Split the last 6 bytes manually because of the stupid guid format.
	last6 := s[len(s)-1]
	if len(last6) != 12 {
		log.Fatalf("wrong number of bytes! expected 12 character suffix, got %v\n", last6)
	}

	s = s[:len(s)-1]
	for i := 0; i < 6; i++ {
		s = append(s, last6[2*i:2*i+2])
	}

	for i := range s {
		n, err := strconv.ParseUint(s[i], 16, bits[i])
		//fmt.Fprintf(os.Stderr, "%08x-", n)
		if err != nil {
			log.Fatalf("%v: not a %d bit Uint: %v", s[i], bits[i], err)
		}
		for bn := 0; bn < bits[i]/8; bn = bn + 1 {
			binary.Write(w, binary.LittleEndian, byte(n))
			n = n >> 8
		}
	}
	//fmt.Fprintf(os.Stderr, "\n")
}

func main() {
	var (
		b    bytes.Buffer
		done bool
	)

	for !done {
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
			writeGUID(&b, g)
		}
		switch op {
		case "TRUE":
			stackdepth++
		case "PUSH":
			stackdepth++
		case "AND":
			stackdepth--
		case "OR":
			stackdepth--
		case "END":
			done = true
		}

	}
	l := b.Len() + 4
	hdr := append([]byte{byte(l), byte(l >> 8), byte(l >> 16), 0x13}, b.Bytes()...)
	if _, err := os.Stdout.Write(hdr); err != nil {
		log.Fatalf("%v", err)
	}
	if stackdepth != 1 {
		log.Fatalf("stackdepth is %d, should be 1", stackdepth)
	}

}
