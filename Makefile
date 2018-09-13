GOFILES_NOVENDOR = $(shell find . -type f -name '*.go' -not -path "./vendor/*" -not -path "./.git/*")

deps:
	go get -u github.com/golang/dep/cmd/dep

format:
	@gofmt -w ${GOFILES_NOVENDOR}
