LOCAL_BIN:=$(CURDIR)/bin
GOBIN?=$(GOPATH)/bin

.PHONY: build
build:
	go build -o $(LOCAL_BIN)/ihp $(CURDIR)/cmd/ihp

.PHONY: install
install:
	go build -o $(GOBIN)/ihp $(CURDIR)/cmd/ihp

.PHONY: lint
lint:
	@golangci-lint run

.gr:
	@rm -fr $(CURDIR)/dist && goreleaser --snapshot

.gb:
	@rm -fr $(CURDIR)/dist && goreleaser build --snapshot