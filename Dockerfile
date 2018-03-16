FROM golang:1.10 as builder

RUN mkdir -p /go/src/github.com/adal-io/vpnswitcher
WORKDIR /go/src/github.com/adal-io/vpnswitcher

COPY . .
COPY .netrc /root

RUN go get -u golang.org/x/vgo 

RUN CC=gcc vgo build -o vpnswitcher

FROM golang:1.10 as app

RUN mkdir -p /go/src/github.com/adal-io/vpnswitcher
WORKDIR /go/src/github.com/adal-io/vpnswitcher

RUN apt-get update && apt-get install -y openvpn sudo

COPY --from=builder /go/src/github.com/adal-io/vpnswitcher/vpnswitcher .

# ENTRYPOINT ["./vpnswitcher"]


