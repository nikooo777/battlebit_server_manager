version := $(shell git describe --abbrev=0 --tags)
commit := $(shell git rev-parse --short HEAD)
commit_long := $(shell git rev-parse HEAD)
branch := $(shell git rev-parse --abbrev-ref HEAD)
curtime := $(shell date "+%Y-%m-%d %T %Z")

BINARY=battlebit_admin_panel
IMPORT_PATH=battlebit_admin_panel
LDFLAGS="-s -w -X ${IMPORT_PATH}/internal/version.version=$(version) -X ${IMPORT_PATH}/internal/version.commit=$(commit) -X ${IMPORT_PATH}/internal/version.commitLong=$(commit_long) -X ${IMPORT_PATH}/internal/version.branch=$(branch) -X '${IMPORT_PATH}/internal/version.date=$(curtime)'"
.DEFAULT_GOAL := linux

.PHONY: test
test:
	go test -cover ./...

.PHONY: lint
lint:
	./scripts/lint.sh

.PHONY: linux
linux:
	GOARCH=amd64 GOOS=linux go build -o dist/linux_amd64/${BINARY} -ldflags ${LDFLAGS}

.PHONY: macos
macos:
	GOARCH=amd64 GOOS=darwin go build -o dist/darwin_amd64/${BINARY} -ldflags ${LDFLAGS}

tag := $(shell git describe --abbrev=0 --tags)
.PHONY: retag
retag:
	@echo "Re-setting tag $(tag) to the current commit"
	git tag -d $(tag)
	git tag $(tag)