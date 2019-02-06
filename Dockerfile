FROM golang:1.11 as builder

WORKDIR /src

COPY . .
COPY .netrc /root

RUN go build -o vpn

FROM golang:1.11 as app

WORKDIR /src

RUN apt-get update && apt-get install -y openvpn sudo

COPY --from=builder /src/vpn .

# ENTRYPOINT ["./vpn"]
