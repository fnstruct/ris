ifneq (,$(wildcard .git))
VERSION ?= $(shell git describe --tags)
else
VERSION ?= 1.0.0
endif

SRC := main.go
BIN := ris

PREFIX ?= /usr/local
BINDIR ?= $(PREFIX)/bin

RISFLAGS ?= -s -w -X main.Version="$(VERSION)"

all: $(BIN)

$(BIN): $(SRC)
	go build -ldflags "$(RISFLAGS)" -trimpath -o $@

.PHONY: install

install:
	install -d $(DESTDIR)$(BINDIR)
	install -m755 $(BIN) $(DESTDIR)$(BINDIR)

.PHONY: clean

clean:
	rm -rf $(BIN)
