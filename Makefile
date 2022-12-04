default: build fmt test readme

build:
	go mod tidy
	go build ./...

fmt:
	go run golang.org/x/tools/cmd/goimports -w .

test:
	go test .

readme:
	go run . -b

day?=$(shell TZ='US/Eastern' date +'%d')

bench:
	go test -benchmem -run=^$$ -bench=Benchmark/Day_$(day) .
