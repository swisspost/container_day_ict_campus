FROM golang:1.24 AS builder
ARG VERSION
WORKDIR /go/src/github.com/the-technat/container-demo/
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /go/helloworld 

FROM scratch
COPY --from=builder /go/helloworld .
ENV GREETING=Hello
ENV ADDR=0.0.0.0:8080
ENTRYPOINT ["/helloworld"]