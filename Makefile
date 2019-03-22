build:
	go build -o rate-crawler cmd/rate-crawler.go
install:
	make build
	mv rate-crawler /usr/local/bin/
.PHONY: clean
clean:
	-rm rate-crawler
