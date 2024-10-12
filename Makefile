SRC := main.go
BIN := ris

PREFIX ?= /usr/local
BINDIR ?= $(PREFIX)/bin

GO_LDFLAGS ?= -s -w

all: $(BIN)

$(BIN): $(SRC)
	go build -ldflags "$(GO_LDFLAGS)" -trimpath -o $@

.PHONY: install

install:
	install -d $(DESTDIR)$(BINDIR)
	install -m755 $(BIN) $(DESTDIR)$(BINDIR)

.PHONY: clean

clean:
	rm -rf $(BIN)
