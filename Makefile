.PHONY: build
build:
	go build -o build/apachedist-resource main.go

test:
	go test ./...

lint:
	golangci-lint run

check-%: build
	cat cmd/testdata/check-$*.json
	cat cmd/testdata/check-$*.json | ./build/apachedist-resource check

in-%: build
	mkdir -p test-output
	cat cmd/testdata/in-$*.json
	cat cmd/testdata/in-$*.json | ./build/apachedist-resource in test-output
	tree test-output
	rm -fr test-output

test-all: test check-tomcat check-tomee check-first-attempt in-tomcat in-tomee
