VERSION=0.9.8
SOURCE?=./...
VINYLDNS_REPO=github.com/vinyldns/vinyldns

all: test build start-api integration stop-api validate-version install

fmt:
	gofmt -s -w vinyldns

test:
	go vet $(SOURCE)
	go test $(SOURCE) -cover

integration:
	go test $(SOURCE) -tags=integration

validate-version:
	cat vinyldns/version.go | grep 'var Version = "$(VERSION)"'

start-api:
	if [ ! -d "$(GOPATH)/src/$(VINYLDNS_REPO)" ]; then \
		echo "$(VINYLDNS_REPO) not found in your GOPATH (necessary for acceptance tests), getting..."; \
		git clone https://$(VINYLDNS_REPO) $(GOPATH)/src/$(VINYLDNS_REPO); \
	fi
	$(GOPATH)/src/$(VINYLDNS_REPO)/bin/docker-up-vinyldns.sh \
		--api-only \
		--version 0.9.1

stop-api:
	./../vinyldns/bin/remove-vinyl-containers.sh

cover:
	go test $(SOURCE) -coverprofile=coverage.out
	go tool cover -html=coverage.out
	rm coverage.out

build:
	go build -ldflags "-X main.version=$(VERSION)" $(SOURCE)

install:
	go install $(SOURCE)

release: deps test
	go get github.com/aktau/github-release
	github-release release \
		--user vinyldns \
		--repo go-vinyldns \
		--tag $(VERSION) \
		--name "$(VERSION)" \
		--description "go-vinyldns version $(VERSION)"
