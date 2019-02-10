build: .netrc
	go build -o vpn

docker: .netrc
	docker build . -t vpn

data:
	rm -rf data
	mkdir data
	curl https://www.privateinternetaccess.com/openvpn/openvpn.zip -o data/openvpn.zip
	cd data && unzip openvpn.zip && rm -f openvpn.zip

.netrc:
	rm -f .netrc
	cp ~/.netrc .netrc

clean:
	rm -f .netrc

.PHONY: clean build data
