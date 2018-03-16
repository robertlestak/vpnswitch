build: .netrc

.netrc:
	rm -f .netrc
	cp ~/.netrc .netrc

clean:
	rm -f .netrc

.PHONY: clean build


