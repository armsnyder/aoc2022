default: fmt test

test:
	go test .

bench:
	go test -run=^$$ -bench .

fmt:
	go run golang.org/x/tools/cmd/goimports -w .

.PHONY: default test fmt
