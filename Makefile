.PHONY: install
install:
	go install

.PHONY: build
build:
	go build -ldflags "-w -s" -o go-graphql-experiments ./main.go

.PHONY: run
run:
	go run main.go
