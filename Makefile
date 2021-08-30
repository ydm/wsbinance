build:
	go build \
		-o bin/${BIN_NAME}

clean:
	rm -fr bin

fix:
	golangci-lint run --fix

lint:
	golangci-lint run

help:
	@echo 'Management commands:'
	@echo
	@echo 'Usage:'
	@echo '    make build           Compile project.'
	@echo '    make clean           Clean directory tree.'
	@echo '    make fix             Fix linting problems.'
	@echo '    make lint            Run linters.'
	@echo '    make help            Print this message.'
	@echo '    make test            Run tests.'
	@echo

test:
	go test ./...

.PHONY: build clean fix lint help test
