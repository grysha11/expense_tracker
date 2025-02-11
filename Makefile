main_path	= ./cmd/api
binary_name	= expense_tracker

build:
	go build -o bin/$(binary_name) $(main_path)

run:
	./bin/$(binary_name)

br:
	$(build) $(run)

tidy:
	go mod tidy

.PHONY: build, run, br, tidy