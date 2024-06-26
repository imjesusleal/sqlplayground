#Variables

APP_NAME = sql-playground
WASM_TARGET = ./static/main.wasm
WASM_OPT = binaryen-version_117-x86_64-linux.tar.gz

#Rules

.PHONY: all
all: front-deps format-deps build-wasm opt-wasm tailwind-dev  

.PHONY: build
build: front-deps format-deps test build-wasm opt-wasm tailwind

.PHONY: lint-deps
lint-deps:
	@command -v golangci-lint >/dev/null 2>&1 || ( \
		curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.58.2; \ )
	golangci-lint run ./src

.PHONY: format-deps
format-deps:
	@command -v gofmt >/dev/null 2>&1 || ( \
		go install golang.org/x/tools/cmd/goimports@latest; \
	) 
	go fmt ./...

.PHONY: opt-wasm
opt-wasm:
	@command -v wasm-opt >/dev/null 2>&1 || ( \
		wget https://github.com/WebAssembly/binaryen/releases/tag/version_117/${WASM_OPT} .; \
		tar -xvzf ${WASM_OPT} .; \
		mv ${WASM_OPT}/bin/wasm-opt /usr/local/bin; \
	)
	wasm-opt ./static/main.wasm --enable-bulk-memory -Oz -o ./static/main.wasm

build-wasm : ./static/main.go
	GOOS=js GOARCH=wasm go build -o ${WASM_TARGET} ./static/main.go

test: 
	GOOS=js GOARCH=wasm go test ./...

serve : 
	go run ./src

front-deps:
	@command -v npm -v >/dev/null 2>&1 || ( \
		curl -fsSL https://deb.nodesource.com/setup_22.x | bash - && \
		apt-get install -y nodejs; \
	)
	npm install .

tailwind:
	npx tailwind -i ./static/styles/input.css -o ./static/styles/output.css

tailwind-dev:
	npx tailwind -i ./static/styles/input.css -o ./static/styles/output.css --watch
