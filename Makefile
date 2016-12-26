.PHONY: all deps test validate lint

all: deps test validate

deps:
	go get -u github.com/golang/lint/golint

test:
	go test -race -cover -coverprofile=coverage.txt ./...

validate: lint
	go vet ./...
	test -z "$(gofmt -s -l . | tee /dev/stderr)"

lint:
	out="$$(golint ./...)"; \
	if [ -n "$$(golint ./...)" ]; then \
		echo "$$out"; \
		exit 1; \
	fi
