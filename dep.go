package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

const (
	glen = 16
	gFmt = " %02x%02x%02x%02x-%02x%02x-%02x%02x-%02x%02x-%02x%02x%02x%02x%02x%02x"
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

	l := int(hdr[0]) + int(hdr[1])*256 + int(hdr[2])*16384 - 4
	//fmt.Fprintf(os.Stderr, "hdr: %v; len is 0x%x bytes\n", hdr, l)

	for {
		if _, err := os.Stdin.Read(op[:]); err != nil {
			if err == io.EOF {
				break
			}
			log.Fatalf("dep file: %v", err)
		}

		l = l - len(op)

		fmt.Printf("%v", opcodes[op[0]])

		if op[0] == 2 {
			g := guid(os.Stdin)
			// This turned out to be the easiest way as opposed to messy for loops,
			// due to the need to print backwards.
			fmt.Printf(gFmt,
				g[3], g[2], g[1], g[0],
				g[5], g[4],
				g[7], g[6],
				g[9], g[8],
				g[10], g[11], g[12], g[13], g[14], g[15])
			l = l - glen
		}

		fmt.Printf("\n")
	}
	if l < 0 {
		log.Fatalf("Tried to read more data than was in the file")
	}
	if l > 0 {
		log.Fatalf("Did not read all the data")
	}
}
