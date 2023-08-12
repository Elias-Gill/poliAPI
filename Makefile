build:
	@rm -r ./build && go build -o build/poliapi

install:
	@go build -o /usr/bin/poliapi && cp -f ./search_links.sh /usr/bin/ && chmod +x /usr/bin/search_links.sh

run:
	@go run main.go

test:
	@gotest ./src/tests
