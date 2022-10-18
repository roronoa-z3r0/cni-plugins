include Makefile.defs

all: usage

usage:
	@echo "usage:"
	@echo  "  \033[35m make build \033[0m:       --- build all plugins"
	@echo  "  \033[35m make image \033[0m:       --- build docker image"
	@echo  "  \033[35m make test \033[0m:        --- run e2e test on your local environment"

.PHONY: build
build:
	@mkdir -p ./.tmp/bin ; \
	for plugin in `ls ./plugins/` ; do   \
		echo "build $${plugin} to $(ROOT_DIR)/.tmp/bin/${plugin}" ; \
		$(GO_BUILD_FLAGS) $(GO_BUILD) -o ./.tmp/bin/$${plugin} ./plugins/$${plugin} ;  \
	done

.PHONY: lint-golang
lint-golang:
	GOOS=linux golangci-lint run ./...

.PHONY: test
test:
	make -C test/test
