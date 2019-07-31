FROM ubuntu:xenial

RUN \
  apt-get update && apt-get upgrade -q -y && \
  apt-get install -y --no-install-recommends golang git make gcc libc-dev ca-certificates && \
  git clone --depth 1 --branch release/1.7 https://github.com/DATx/go-DATx && \
  (cd go-DATx && make gdatx) && \
  cp go-DATx/build/bin/gdatx /gdatx && \
  apt-get remove -y golang git make gcc libc-dev && apt autoremove -y && apt-get clean && \
  rm -rf /go-DATx

EXPOSE 8545
EXPOSE 30303

ENTRYPOINT ["/gdatx"]
