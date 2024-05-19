main.wasm : main.go
	GOOS=js GOARCH=wasm go build -o ./static/main.wasm ./static/main.go

