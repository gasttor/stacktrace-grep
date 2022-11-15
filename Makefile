BINARY_NAME=sgrep

build:
	go build -o bin/${BINARY_NAME} src/sgrep.go

clean:
	go clean
	rm bin/${BINARY_NAME}
