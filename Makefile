BUILD_LINUX := go build -ldflags="-s -w"

all:
	@go fmt
	@mkdir -p output
	@$(BUILD_LINUX) -o output/main

clean:
	@rm -rf output/
