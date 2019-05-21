
all: mod fmt test build

clean:
	rm -rf bin
	go clean

install: 
	go install -v

test:
	go test ./...

fmt:
	go fmt ./... 

mod:
	go mod download

image:
	docker build -t imagename:latest .

release:
	echo "nothing todo"

build:
	mkdir -p bin
	go build -o ./bin/main *.go 

.PHONY: clean install test fmt image release build mod