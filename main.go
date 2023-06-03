package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/ToshihitoKon/codename-generator/utils"
	v1 "github.com/ToshihitoKon/codename-generator/v1"
)

func main() {
	var text string
	var generator utils.CodenameGenerator

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

	// Generator version
	generator = v1.New()

	codename, err := generator.GenerateCodename(text)
	if err != nil {
		fmt.Printf("error: %s", err)
		os.Exit(1)
	}

	fmt.Printf(codename)
}
