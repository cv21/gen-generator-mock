test:
	go test ./...

install:
	go build -o ${GOPATH}/pkg/gen/generator/github.com/cv21/gen-generator-mock@v1.0.0/generator ./main.go