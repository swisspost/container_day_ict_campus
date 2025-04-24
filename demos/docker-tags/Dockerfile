FROM golang:1.24 AS builder
ARG VERSION
WORKDIR /go/src/github.com/the-technat/container-demo/
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /go/helloworld -ldflags="-X 'main.Version=${VERSION}'"

FROM scratch
COPY --from=builder /go/helloworld .
ENTRYPOINT ["/helloworld"]
CMD ["-addr=0.0.0.0:8080"]