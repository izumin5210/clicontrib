.DEFAULT_GOAL := all

ORG := github.com/izumin5210
PROJECT := clicontrib
ROOT_PKG ?= $(ORG)/$(PROJECT)

SRC_FILES := $(shell git ls-files --cached --others --exclude-standard | grep -E "\.go$$")
GOFMT_TARGET := $(SRC_FILES)
GOLINT_TARGET := $(shell go list ./...)

XC_ARCH := 386 amd64
XC_OS := darwin linux windows

#  Utils
#----------------------------------------------------------------
define section
  @printf "\e[34m--> $1\e[0m\n"
endef
 
#  dep
#----------------------------------------------------------------
DEP_BIN_DIR := ./vendor/.bin/
DEP_SRCS := \
	github.com/mitchellh/gox

DEP_BINS := $(addprefix $(DEP_BIN_DIR),$(notdir $(DEP_SRCS)))

define dep-bin-tmpl
$(eval OUT := $(addprefix $(DEP_BIN_DIR),$(notdir $(1))))
$(OUT): dep
	$(call section,Installing $(OUT))
	@cd vendor/$(1) && GOBIN="$(shell pwd)/$(DEP_BIN_DIR)" go install .
endef

$(foreach src,$(DEP_SRCS),$(eval $(call dep-bin-tmpl,$(src))))

#  App
#----------------------------------------------------------------
BIN_DIR := ./bin/
OUT_DIR := ./dist
GENERATED_BINS :=
PACKAGES :=
CMDS := $(wildcard ./cmd/*)

LDFLAGS_CMD = ./bin/cliutils --config cliutils-$1.toml ldflags

define cmd-tmpl
$(eval NAME := $(notdir $(1)))
$(eval OUT := $(addprefix $(BIN_DIR),$(NAME)))
$(eval GENERATED_BINS += $(OUT))
$(OUT): $(SRC_FILES)
	$(call section,Building $(OUT))
	go build $(GO_BUILD_FLAGS) -o $(OUT) $(1)
	go build $(GO_BUILD_FLAGS) -ldflags "$$$$($(call LDFLAGS_CMD,$(NAME)))" -o $(OUT) $(1)

.PHONY: $(NAME)
$(NAME): $(OUT)

$(eval PACKAGES += $(NAME)-package)

.PHONY: $(NAME)-package
$(NAME)-package: $(NAME)
	@PATH=$(shell pwd)/$(DEP_BIN_DIR):$$$$PATH gox \
		-ldflags="$$$$($(call LDFLAGS_CMD,$(NAME)))" \
		-os="$(XC_OS)" \
		-arch="$(XC_ARCH)" \
		-output="$(OUT_DIR)/$(NAME)_{{.OS}}_{{.Arch}}" \
		$(1)
endef

$(foreach src,$(CMDS),$(eval $(call cmd-tmpl,$(src))))

.PHONY: all
all: $(GENERATED_BINS)

#  Commands
#----------------------------------------------------------------
.PHONY: setup
setup: dep $(DEP_BINS)

.PHONY: clean
clean:
	rm -rf $(BIN_DIR)/*

.PHONY: clobber
clobber: clean
	rm -rf vendor

.PHONY: dep
dep: Gopkg.toml Gopkg.lock
	$(call section,Installing dependencies)
	@dep ensure -v -vendor-only

.PHONY: gen
gen:
	@PATH=$(shell pwd)/$(DEP_BIN_DIR):$$PATH go generate ./...

.PHONY: lint
lint:
	$(call section,Linting)
	@gofmt -e -d -s $(GOFMT_TARGET) | awk '{ e = 1; print $0 } END { if (e) exit(1) }'
	@echo $(GOLINT_TARGET) | xargs -n1 golint -set_exit_status

.PHONY: test
test:
	$(call section,Testing)
	@go test $(GO_TEST_FLAGS) ./...

.PHONY: cover
cover:
	$(call section,Testing with coverage)
	@go test $(GO_TEST_FLAGS) $(GO_COVER_FLAGS) ./...

.PHONY: packages
packages: $(PACKAGES)
