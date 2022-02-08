FROM golang:latest
# FROM golang:1.17.6

RUN apt update \
    && apt install -y --no-install-recommends \
    apt-utils \
    git \
    ca-certificates \
    procps \
    && apt autoremove -y \
    && apt clean -y \
    && rm -rf /var/lib/apt/lists/* \
    && go install github.com/go-delve/delve/cmd/dlv@latest \
    && go install honnef.co/go/tools/cmd/staticcheck@latest \
    && go get -v github.com/ramya-rao-a/go-outline \
    && go get -v golang.org/x/tools/gopls
