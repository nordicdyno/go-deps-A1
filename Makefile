.PHONY: all
all: deps test

.PHONY: deps
deps:
	./tools/go-get-private.sh
	go mod tidy

.PHONY: test
test:
	go test -v -count=1 ./...