SHELL=bash
VERSION=0.9.16
SOURCE?=./...
VINYLDNS_VERSION=0.10.4

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

.PHONY: start-api
start-api: stop-api
	@set -euo pipefail
	echo "Starting VinylDNS API.."
	docker run -d --name vinyldns-go-api -p "9000:9000" -p "19001:19001" -e RUN_SERVICES="all tail-logs" vinyldns/build:base-test-integration-v$(VINYLDNS_VERSION)
	echo "Waiting for VinylDNS API to start.."
	{ timeout "20s" grep -q 'STARTED SUCCESSFULLY' <(timeout 20s docker logs -f vinyldns-go-api) ;	echo "VinylDNS API STARTED SUCCESSFULLY";  } || { echo "VinylDNS API failed to start"; exit 1; }

.PHONY: stop-api
stop-api:
	@set -euo pipefail
	if docker ps | grep -q "vinyldns-go-api"; then
		docker kill vinyldns-go-api &> /dev/null || true
		docker rm vinyldns-go-api &> /dev/null || true   
	fi

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