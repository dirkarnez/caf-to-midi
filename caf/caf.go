package caf

import (
	"bytes"
	"fmt"
	"log"

	"github.com/pascoej/caf"
)

func stringToChunkType(str string) (result caf.FourByteString) {
	for i, v := range str {
		result[i] = byte(v)
	}
	return
}

func ConvertCafToMidi(cafBytes []byte) ([]byte, error) {
	reader := bytes.NewReader(cafBytes)
	f := &caf.File{}
	if err := f.Decode(reader); err != nil {
		log.Fatal(err)
	}

	for _, chunk := range f.Chunks {
		if chunk.Header.ChunkType == stringToChunkType("midi") {
			b, ok := chunk.Contents.([]byte)
			if !ok {
				return nil, fmt.Errorf("not bytes")
			} else {
				return b, nil
			}
		}
	}
	return nil, fmt.Errorf("fail to convert")
}
