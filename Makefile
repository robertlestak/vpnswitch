build: .netrc
	go build -o vpn

docker: .netrc
	docker build . -t vpn

.netrc:
	rm -f .netrc
	cp ~/.netrc .netrc

clean:
	rm -f .netrc

.PHONY: clean build
