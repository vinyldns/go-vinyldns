VERSION=0.8.0
SOURCE?=./...

all: deps test install

deps:
	@go tool cover 2>/dev/null; if [ $$? -eq 3 ]; then \
		go get -u golang.org/x/tools/cmd/cover; \
	fi
	go get -u github.com/golang/dep/cmd/dep
	dep ensure

test:
	go vet $(SOURCE)
	go test $(SOURCE) -cover

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
