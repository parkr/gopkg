GOPKG=github.com/parkr/gopkg

all: build test

build:
	go install $(GOPKG)/...

test:
	go test $(GOPKG)/...
