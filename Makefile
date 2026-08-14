# machina -- machine layout for agents
# suckless-style Makefile; no hidden magic.

PREFIX    ?= $(HOME)/.local
BIN        = machina
VERSION    = 0.1.0

GOFLAGS   ?= -trimpath
LDFLAGS   ?= -s -w -X main.version=$(VERSION)

CMD = ./cmd/machina
SRC = $(wildcard cmd/machina/*.go) $(wildcard scaffold/*.go) $(wildcard skills/*.go) $(wildcard skills/*/SKILL.md)

all: $(BIN)

$(BIN): $(SRC) go.mod
	CGO_ENABLED=0 go build $(GOFLAGS) -ldflags "$(LDFLAGS)" -o $(BIN) $(CMD)

tidy:
	go mod tidy

vet:
	go vet ./...

test:
	go test ./...

fmt:
	gofmt -w cmd/machina scaffold

clean:
	rm -f $(BIN)

install: $(BIN)
	install -Dm755 $(BIN) $(DESTDIR)$(PREFIX)/bin/$(BIN)
	ln -sf $(BIN) $(DESTDIR)$(PREFIX)/bin/disk
	ln -sf $(BIN) $(DESTDIR)$(PREFIX)/bin/mem
	ln -sf $(BIN) $(DESTDIR)$(PREFIX)/bin/load
	ln -sf $(BIN) $(DESTDIR)$(PREFIX)/bin/home-map

uninstall:
	rm -f $(DESTDIR)$(PREFIX)/bin/$(BIN)
	rm -f $(DESTDIR)$(PREFIX)/bin/disk
	rm -f $(DESTDIR)$(PREFIX)/bin/mem
	rm -f $(DESTDIR)$(PREFIX)/bin/load
	rm -f $(DESTDIR)$(PREFIX)/bin/home-map

.PHONY: all tidy vet test fmt clean install uninstall
