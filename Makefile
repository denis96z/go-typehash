.PHONY: all
all: typehash-gen

.PHONY: typehash-gen
typehash-gen:
	go build -o $(PWD)/bin/typehash-gen $(PWD)/cmd/typehash-gen/...

.PHONY: fmt
fmt:
	go fmt $(PWD)/...
