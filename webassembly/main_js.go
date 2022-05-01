//go:build js && wasm
// +build js,wasm

package main

//go:generate cp $GOROOT/misc/wasm/wasm_exec.js .

import (
	"path/filepath"
	"strings"
	"syscall/js"

	dirkCaf "github.com/dirkarnez/caf-to-midi/caf"
)

func GetFileName(path string) string {
	noExtension := strings.TrimSuffix(path, filepath.Ext(path))
	return noExtension[strings.LastIndex(noExtension, "\\")+1:]
}

func convertCafToMidi_JavaScript(this js.Value, args []js.Value) interface{} {
	buffer := make([]byte, args[0].Length())

	js.CopyBytesToGo(buffer, args[0])

	// 计算md5的值
	bytes, _ := dirkCaf.ConvertCafToMidi(buffer)

	dst := js.Global().Get("Uint8Array").New(len(bytes))
	js.CopyBytesToJS(dst, bytes)

	return dst
}

func main() {
	js.Global().Set("convertCafToMidiGo", js.FuncOf(convertCafToMidi_JavaScript))
	select {} // block the main thread forever
}
