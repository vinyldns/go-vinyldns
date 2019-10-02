VERSION=0.9.9
SOURCE?=./...
VINYLDNS_REPO=github.com/vinyldns/vinyldns
VINYLDNS_VERSION=0.9.3

all: test build start-api integration stop-api validate-version install

fmt:
	gofmt -s -w vinyldns

test:
	go vet $(SOURCE)
	GO111MODULE=on go test $(SOURCE) -cover

integration:
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

release: test
	go get github.com/aktau/github-release
	github-release release \
		--user vinyldns \
		--repo go-vinyldns \
		--tag $(VERSION) \
		--name "$(VERSION)" \
		--description "go-vinyldns version $(VERSION)"
