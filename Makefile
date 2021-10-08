VERSION=0.9.15
SOURCE?=./...
VINYLDNS_REPO=github.com/vinyldns/vinyldns
VINYLDNS_VERSION=0.9.10

ifndef $(GOPATH)
    GOPATH=$(shell go env GOPATH)
    export GOPATH
endif

all: check-fmt test build integration stop-api validate-version install

fmt:
	gofmt -s -w vinyldns

check-fmt:
	test -z "$(shell gofmt -s -l vinyldns | tee /dev/stderr)"

test:
	go vet $(SOURCE)
	GO111MODULE=on go test $(SOURCE) -cover

integration: start-api
	GO111MODULE=on go test $(SOURCE) -tags=integration

validate-version:
	cat vinyldns/version.go | grep 'var Version = "$(VERSION)"'

start-api:
	if [ ! -d "$(GOPATH)/src/$(VINYLDNS_REPO)-$(VINYLDNS_VERSION)" ]; then \
		echo "$(VINYLDNS_REPO)-$(VINYLDNS_VERSION) not found in your GOPATH (necessary for acceptance tests), getting..."; \
		git clone \
			--branch v$(VINYLDNS_VERSION) \
			https://$(VINYLDNS_REPO) \
			$(GOPATH)/src/$(VINYLDNS_REPO)-$(VINYLDNS_VERSION); \
	fi
	$(GOPATH)/src/$(VINYLDNS_REPO)-$(VINYLDNS_VERSION)/bin/docker-up-vinyldns.sh \
		--api-only \
		--version $(VINYLDNS_VERSION)

stop-api:
	$(GOPATH)/src/$(VINYLDNS_REPO)-$(VINYLDNS_VERSION)/bin/remove-vinyl-containers.sh

build:
	GO111MODULE=on go build -ldflags "-X main.version=$(VERSION)" $(SOURCE)

install:
	GO111MODULE=on go install $(SOURCE)

release: test validate-version
	go get github.com/aktau/github-release
	github-release release \
		--user vinyldns \
		--repo go-vinyldns \
		--tag $(VERSION) \
		--name "$(VERSION)" \
		--description "go-vinyldns version $(VERSION)"
