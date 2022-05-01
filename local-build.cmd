SET GOOS=windows
SET GOARCH=amd64
go build

SET GOOS=js
SET GOARCH=wasm

go build -o build-webassembly\main.wasm  .\webassembly\main_js.go

copy "%GOROOT%misc\wasm\wasm_exec.js" build-webassembly\