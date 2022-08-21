FROM ubuntu:22.04

ARG GOLANG_VERSION=1.19

RUN apt-get update
RUN apt-get install -y git jq wget zip gcc tar

RUN wget https://go.dev/dl/go${GOLANG_VERSION}.linux-amd64.tar.gz
RUN tar -xf go${GOLANG_VERSION}.linux-amd64.tar.gz
RUN mv go /usr/bin
RUN rm go${GOLANG_VERSION}.linux-amd64.tar.gz
ENV PATH="/usr/bin/go/bin:$PATH"

RUN go install github.com/uudashr/gopkgs/v2/cmd/gopkgs@latest
RUN go install github.com/ramya-rao-a/go-outline@latest
RUN go install github.com/cweill/gotests/gotests@latest
RUN go install github.com/fatih/gomodifytags@latest
RUN go install github.com/josharian/impl@latest
RUN go install github.com/haya14busa/goplay/cmd/goplay@latest
RUN go install github.com/go-delve/delve/cmd/dlv@latest
RUN mv /root/go/bin/dlv /root/go/bin/dlv-dap
RUN go install github.com/go-delve/delve/cmd/dlv@latest
RUN go install honnef.co/go/tools/cmd/staticcheck@latest
RUN go install golang.org/x/tools/gopls@latest
