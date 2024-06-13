SHELL=bash
VERSION=0.9.16
SOURCE?=./...
VINYLDNS_REPO=github.com/vinyldns/vinyldns
VINYLDNS_DIR="$(GOPATH)/src/$(VINYLDNS_REPO)/" 
VINYLDNS_VERSION=latest

# Check that the required version of make is being used
REQ_MAKE_VER:=3.82
ifneq ($(REQ_MAKE_VER),$(firstword $(sort $(MAKE_VERSION) $(REQ_MAKE_VER))))
   $(error The version of MAKE $(REQ_MAKE_VER) or higher is required; you are running $(MAKE_VERSION))
endif

ifndef $(GOPATH)
    GOPATH=$(shell go env GOPATH)
    export GOPATH
endif
.ONESHELL:

all: check-fmt test build integration stop-api validate-version install

.PHONY: fmt
fmt:
	gofmt -s -w vinyldns

.PHONY: check-fmt
check-fmt:
	test -z "$(shell gofmt -s -l vinyldns | tee /dev/stderr)"

.PHONY: all-test
all-test: test integration

.PHONY: test
test:
	go vet $(SOURCE)
	GO111MODULE=on go test $(SOURCE) -cover

.PHONY: integration
integration: start-api
	GO111MODULE=on go test $(SOURCE) -tags=integration

.PHONY: validate-version
validate-version:
	cat vinyldns/version.go | grep 'var Version = "$(VERSION)"'

.PHONY: clone-vinyl
clone-vinyl:
	if [ ! -d  $(VINYLDNS_DIR) ]; then \
		echo "$(VINYLDNS_REPO) not found in your GOPATH (necessary for acceptance tests), getting..."; \
		git clone \
			https://$(VINYLDNS_REPO) \
			$(VINYLDNS_DIR); \
	else \
		git -C $(VINYLDNS_DIR) pull ; \
	fi
  
.PHONY: start-api
start-api: clone-vinyl stop-api
	$(GOPATH)/src/$(VINYLDNS_REPO)/quickstart/quickstart-vinyldns.sh \
		--api --version-tag $(VINYLDNS_VERSION)

.PHONY: stop-api
stop-api:
	$(GOPATH)/src/$(VINYLDNS_REPO)/quickstart/quickstart-vinyldns.sh \
		--clean

.PHONY: build
build:
	GO111MODULE=on go build -ldflags "-X main.version=$(VERSION)" $(SOURCE)

.PHONY: install
install:
	GO111MODULE=on go install $(SOURCE)

.PHONY: release
release: test validate-version
	go get github.com/aktau/github-release
	github-release release \
		--user vinyldns \
		--repo go-vinyldns \
		--tag $(VERSION) \
		--name "$(VERSION)" \
		--description "go-vinyldns version $(VERSION)"