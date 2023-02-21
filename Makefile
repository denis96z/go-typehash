.PHONY: all
all: typehash

.PHONY: typehash-gen
typehash:
	go build -o $(PWD)/bin/typehash-gen $(PWD)/cmd/typehash-gen/...

.PHONY: fmt
fmt:
	go fmt $(PWD)/...

.PHONY: tidy
tidy:
	go mod tidy
