FROM golang:1 AS build

WORKDIR "$GOPATH/src/github.com/Nekroze/chell"

# Tools
RUN curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | bash -s -- -b "$GOPATH/bin" v1.14.0
RUN go get -u github.com/kyoh86/richgo
# Go 1.11+ modules
ENV GO111MODULE=on

# Deps
COPY go.* ./
RUN go mod download

# Copy in source code
COPY main.go ./
COPY ./pkg ./pkg
# Test lint and compile
RUN golangci-lint run --skip-files '(parser|lexer)\.(go|y|rl)' --deadline '2m' --enable-all --disable gochecknoglobals,gochecknoinits,gocyclo
RUN richgo test -v ./...
ENV CGO_ENABLED=0 GOOS=linux GOARCH=386
RUN go build \
    -a -installsuffix cgo -ldflags='-w -s' -o /usr/bin/chell \
    .


FROM nekroze/containaruba:alpine AS test
CMD ["--order=random"]

RUN apk add --no-cache tmux

COPY --from=build /usr/bin/chell /usr/bin/
COPY ./features /usr/src/app/features


FROM alpine AS final

RUN apk add --no-cache tmux

COPY --from=build /usr/bin/chell /usr/bin/chell
ENTRYPOINT ["/usr/bin/chell"]
