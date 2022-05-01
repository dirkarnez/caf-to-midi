package main

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	dirkCaf "github.com/dirkarnez/caf-to-midi/caf"
)

func GetFileName(path string) string {
	noExtension := strings.TrimSuffix(path, filepath.Ext(path))
	return noExtension[strings.LastIndex(noExtension, "\\")+1:]
}

func convertCafToMidi_Desktop(filePath string) {
	contents, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}

	out, err := dirkCaf.ConvertCafToMidi(contents)
	if err != nil {
		log.Fatal(err)
		return
	}

	file, err := os.Create(GetFileName(filePath) + ".mid")
	if err != nil {
		log.Fatal(err)
		return
	}
	defer file.Close()
	file.Write(out)
}

func main() {
	if len(os.Args) == 2 {
		convertCafToMidi_Desktop(os.Args[1:2][0])
	} else {
		log.Fatal("please drag a file")
	}
}
