.PHONY: example-% full-% test test-%

example-%:
	go run cmd/$*/main.go cmd/$*/input/example.txt

full-%:
	go run cmd/$*/main.go cmd/$*/input/full.txt

test:
	go test ./internal/...

test-%:
	go test ./internal/$*/...

