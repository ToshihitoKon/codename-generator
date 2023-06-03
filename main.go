package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	flag "github.com/spf13/pflag"

	"github.com/ToshihitoKon/codename-generator/utils"
	v1 "github.com/ToshihitoKon/codename-generator/v1"
)

var GeneratorVersions = map[int]utils.CodenameGenerator{
	1: v1.New(),
}

func printAvailableVersions() {
	var versions []int
	for v, _ := range GeneratorVersions {
		versions = append(versions, v)
	}
	fmt.Fprintln(os.Stderr, "available generator versions:", versions)
}

func main() {
	var text string
	var generator utils.CodenameGenerator

	var optGeneratorVersion *int = flag.IntP("generator-version", "g", -1, "Generator Version")
	var optList *bool = flag.BoolP("list", "l", false, "List generator versions")

	flag.Parse()

	if *optList {
		printAvailableVersions()
		os.Exit(0)
	}

	generator, ok := GeneratorVersions[*optGeneratorVersion]
	if !ok {
		if *optGeneratorVersion < 0 {
			fmt.Fprintln(os.Stderr, "error: --generator-version is required")
		} else {
			fmt.Fprintf(os.Stderr, "error: invalid Generator Version: %d\n", *optGeneratorVersion)
		}
		printAvailableVersions()
		os.Exit(1)
	}

	args := flag.Args()
	if 0 < len(args) {
		text = args[0]
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

	codename, err := generator.GenerateCodename(text)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s", err)
		os.Exit(1)
	}

	fmt.Printf(codename)
}
