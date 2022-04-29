package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/pascoej/caf"
)

func stringToChunkType(str string) (result caf.FourByteString) {
	for i, v := range str {
		result[i] = byte(v)
	}
	return
}

func GetFileName(path string) string {
	noExtension := strings.TrimSuffix(path, filepath.Ext(path))
	return noExtension[strings.LastIndex(noExtension, "\\")+1:]
}

func main() {
	args := os.Args[1:2]
	if len(args) != 1 {
		log.Fatal("Please drag a file")
	}

	filePath := args[0]
	contents, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}

	reader := bytes.NewReader(contents)
	f := &caf.File{}
	if err := f.Decode(reader); err != nil {
		log.Fatal(err)
	}

	for _, chunk := range f.Chunks {
		if chunk.Header.ChunkType == stringToChunkType("midi") {
			b, ok := chunk.Contents.([]byte)
			if !ok {
				log.Fatal("Not bytes")
			}

			file, err := os.Create(GetFileName(filePath) + ".mid")
			if err != nil {
				log.Fatal(err)
			}
			defer file.Close()
			file.Write(b)
			break
		}
	}
}
