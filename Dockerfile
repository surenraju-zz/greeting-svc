FROM golang AS builder
MAINTAINER suren.1988@gmail.com
ENV GOBIN ${GOPATH}/bin
ENV PATH $PATH:$GOBIN:/usr/local/include/bin:$GOROOT/bin
RUN apt-get update -y && apt-get install unzip tree
WORKDIR /go/src/github.com/
COPY greeting-svc ./greeting-svc
RUN mv ./greeting-svc/go.mod ./greeting-svc/go.sum .
RUN export CGO_ENABLED=0; go build ./greeting-svc/server

FROM alpine 
WORKDIR /usr/app
COPY --from=builder /go/src/github.com/server .
ENV PATH $PATH:/usr/app
CMD ["server"]  