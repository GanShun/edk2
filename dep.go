package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

const (
	glen = 16
)

var (
	opcodes = map[byte]string{
		6: "TRUE",
		2: "PUSH",
		3: "AND",
		4: "OR",
		8: "END",
	}
)

func guid(r io.Reader) []byte {
	var b [16]byte
	if n, err := r.Read(b[:]); n != 16 || err != nil {
		log.Fatalf("Reading GUID: got %d of 16 bytes: %v", n, err)
	}
	return b[:]
}

func main() {
	var (
		hdr [4]byte
		op  [1]byte
	)

	if n, err := os.Stdin.Read(hdr[:]); n != len(hdr) || err != nil {
		log.Fatalf("Reading header: %v", err)
	}

	if hdr[3] != 0x13 {
		log.Fatalf("Header type is 0x%x, not 0x13", hdr[3])
	}

	len := int(hdr[0]) + int(hdr[1])*256 + int(hdr[2])*16384
	fmt.Printf("hdr: %v; Len is 0x%x bytes\n", hdr, len)

	for {
		if _, err := os.Stdin.Read(op[:]); err != nil {
			log.Fatalf("dep file: %v", err)
		}

		fmt.Printf("%v", opcodes[op[0]])

		if op[0] == 2 {
			g := guid(os.Stdin)
			fmt.Printf(" %v", g)
		}
		fmt.Printf("\n")
	}
}
