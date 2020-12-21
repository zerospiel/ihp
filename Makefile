LOCAL_BIN:=$(CURDIR)/bin
ifeq ($(GOBIN),)
GOBIN:=$(GOPATH)/bin
endif

.PHONY: build
build:
	go build -o $(LOCAL_BIN)/ihp $(CURDIR)/cmd/ihp

.PHONY: install
install:
	go build -o $(GOBIN)/ihp $(CURDIR)/cmd/ihp