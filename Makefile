.DEFAUKLT_GOAL := test
.PHONY: FORCE

# enable consistent Go 1.12/1.13 GOPROXY behavior.
export GOPROXY = https://proxy.golang.org

BINARY = github-templater
OUTPUT_DIR = ./out/binaries
OUTPUT_BINARY = $(OUTPUT_DIR)/$(BINARY)

ifeq ($(OS),Windows_NT)
	BINARY := $(BINARY).exe
	OUTPUT_DIR := .\out\binaries
	OUTPUT_BINARY := $(OUTPUT_DIR)\$(BINARY)
endif

# Build
build: $(OUTPUT_BINARY)
.PHONY: build

build_race:
	go build -race -o $(OUTPUT_BINARY) ./cmd/github-templater
.PHONY: build_race

clean:
	@echo Cleaning up...
ifneq ("$(wildcard $(OUTPUT_DIR)/$(OUTPUT_BINARY))","")
	rm -f $(OUTPUT_BINARY)
endif
ifneq ("$(wildcard tools/goreleaser)","")
	rm -f tools/goreleaser
endif
	@echo Done.
.PHONY: clean

# Test
test: export DEBUG_MODE = true
test: build
	GL_TEST_RUN=1 go test -v -parallel 2 ./...

test_race: build_race
	GL_TEST_RUN=1 ./$(OUTPUT_BINARY) run -v --timeout=5m
.PHONY: test_race

fast_generate: assets/github-action-config.json
.PHONY: fast_generate

fast_check_generated:
	$(MAKE) --always-make fast_generate
	git checkout -- go.mod go.sum # can differ between go1.16 and go1.17
	git diff --exit-code # check no changes

release: .goreleaser.yml tools/goreleaser
	./tools/goreleaser
.PHONY: release

snapshot: .goreleaser.yml tools/goreleaser
	./tools/goreleaser --snapshot --rm-dist
.PHONY: snapshot

# Non-PHONY targets (real files)

$(OUTPUT_BINARY): FORCE
	@echo Building $(OUTPUT_BINARY)...
	go build -o $@ ./cmd/github-templater

tools/goreleaser: export GOFLAGS = -mod=readonly
tools/goreleaser: tools/go.mod tools/go.sum
	cd tools && go build github.com/goreleaser/goreleaser

# TODO: migrate to docs/
tools/Dracula.itermcolors:
	curl -fL -o $@ https://raw.githubusercontent.com/dracula/iterm/master/Dracula.itermcolors

assets/github-action-config.json: FORCE $(BINARY)
	# go run ./scripts/gen_github_action_config/main.go $@
	cd ./scripts/gen_github_action_config/; go run ./main.go ../../$@

go.mod: FORCE
	go mod tidy
	go mod verify
go.sum: go.mod