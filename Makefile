VERSION=0.9.0
SOURCE?=./...
VINYLDNS_REPO=github.com/vinyldns/vinyldns

all: deps test start-api integration stop-api install

deps:
	@go tool cover 2>/dev/null; if [ $$? -eq 3 ]; then \
		go get -u golang.org/x/tools/cmd/cover; \
	fi
	go get -u github.com/golang/dep/cmd/dep
	dep ensure

fmt:
	gofmt -s -w vinyldns

test:
	go vet $(SOURCE)
	go test $(SOURCE) -cover

integration:
	go test $(SOURCE) -tags=integration

start-api:
	if [ ! -d "$(GOPATH)/src/$(VINYLDNS_REPO)" ]; then \
		echo "$(VINYLDNS_REPO) not found in your GOPATH (necessary for acceptance tests), getting..."; \
		git clone https://$(VINYLDNS_REPO) $(GOPATH)/src/$(VINYLDNS_REPO); \
	fi
	$(GOPATH)/src/$(VINYLDNS_REPO)/bin/docker-up-vinyldns.sh \
		--api-only \
		--version 0.8.0

stop-api:
	./../vinyldns/bin/remove-vinyl-containers.sh

cover:
	go test $(SOURCE) -coverprofile=coverage.out
	go tool cover -html=coverage.out
	rm coverage.out

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
