package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"

	cgenutils "github.com/ToshihitoKon/codename-generator/utils"
	cgenv1 "github.com/ToshihitoKon/codename-generator/v1"
)

var generatorVersions = map[int]cgenutils.CodenameGenerator{
	1: cgenv1.New(),
}

func main() {
	run()
}

func run() {
	http.HandleFunc("/", handler)
	if err := http.ListenAndServe(":"+os.Getenv("PORT"), nil); err != nil {
		log.Println(err.Error())
		return
	}

}

func handler(w http.ResponseWriter, req *http.Request) {
	genver := req.FormValue("generator-version")
	text := req.FormValue("text")
	if genver == "" || text == "" {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, "generator-version and text parameters required")
		return
	}
	version, err := strconv.ParseInt(genver, 10, 0)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, fmt.Sprintf("invalid generator-version %s ", genver))
		return
	}

	gen, ok := generatorVersions[int(version)]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, fmt.Sprintf("invalid generator-version %s ", genver))
		return
	}

	codename, err := gen.GenerateCodename(text)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	io.WriteString(w, codename)
}
