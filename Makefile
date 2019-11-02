BINARY_PATH=bin
COUNTERS_API_PATH=cmd/counters-api/main.go
COUNTERS_API_BIN=$(BINARY_PATH)/counters-api

build:
	go build -o $(COUNTERS_API_BIN) -race -v $(COUNTERS_API_PATH)

run: build
	./$(COUNTERS_API_BIN)

test:
	gotest -race -v -count=1 -timeout=10s ./...

clean:
	go clean $(COUNTERS_API_PATH)
	rm -f $(BINARY_PATH)/*