package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"os"

	"github.com/pascoej/caf"
)

func stringToChunkType(str string) (result caf.FourByteString) {
	for i, v := range str {
		result[i] = byte(v)
	}
	return
}

func main() {
	contents, err := ioutil.ReadFile("./Delicate Piano 18.caf")
	if err != nil {
		log.Fatal(err)
	}
	if len(contents) == 0 {
		log.Fatal("testing with empty file")
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

			file, err := os.Create("output.midi")
			if err != nil {
				log.Fatal(err)
			}
			defer file.Close()
			file.Write(b)
			break
		}
	}
}
