FROM alpine:3.5

RUN \
  apk add --update go git make gcc musl-dev linux-headers ca-certificates && \
  git clone --depth 1 --branch release/1.7 https://github.com/KunkaYU/go-DATx && \
  (cd go-datx && make gdatx) && \
  cp go-datx/build/bin/gdatx /gdatx && \
  apk del go git make gcc musl-dev linux-headers && \
  rm -rf /go-datx && rm -rf /var/cache/apk/*

EXPOSE 8545
EXPOSE 30303

ENTRYPOINT ["/gdatx"]
