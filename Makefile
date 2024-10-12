SRC := main.go
BIN := ris

GO_LDFLAGS ?= -s -w

all: $(BIN)

$(BIN): $(SRC)
	go build -ldflags "$(GO_LDFLAGS)" -trimpath -o $@

.PHONY: clean

clean:
	rm -rf $(BIN)
