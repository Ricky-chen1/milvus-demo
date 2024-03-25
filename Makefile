DIR = $(shell pwd)
CMD = $(DIR)/cmd

# options: amd64 arm64
ARCH = arm64

# options: linux darwin windows
OS = linux

# define the bin path
BIN = $(DIR)/bin

# go settings
GO = go
GO_BUILD = $(GO) build
GO_BUILD_FLAGS = -v
GO_BUILD_LDFLAGS = -X main.version=$(VERSION)

.PHONY: env-up
env-up:
	docker compose up -d

.PHONY: env-down
env-down:
	docker compose down

.PHONY: clean
clean:
	sudo rm -rf ./volumes