package main

import (
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"strings"
)

func main() {
	var text string
	if 1 < len(os.Args) {
		text = os.Args[1]
	} else {
		f := os.Stdin
		buffer, err := io.ReadAll(f)
		if err != nil {
			log.Printf("error: %s", err)
			os.Exit(1)
		}
		text = string(buffer)
		text = strings.TrimSuffix(text, "\n")
	}

	textHashBytes := sha256.Sum256([]byte(text))
	randSourceInt64, readBytes := binary.Varint(textHashBytes[:])
	if readBytes < 0 {
		log.Printf("error: overflow")
		os.Exit(1)
	}

	r := rand.New(rand.NewSource(randSourceInt64))
	fmt.Printf("%s %s", adjectives[r.Intn(len(adjectives))], sweets[r.Intn(len(sweets))])
}
