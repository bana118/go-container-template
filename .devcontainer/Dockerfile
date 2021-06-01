FROM golang:1.16.4-buster

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
    && curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.40.1
