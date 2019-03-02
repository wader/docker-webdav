FROM golang:1.12 AS builder
RUN mkdir -p /go/src/app
WORKDIR /go/src/app
COPY *.go /go/src/app
RUN go get && go install -installsuffix netgo -tags netgo

FROM scratch
COPY --from=builder /go/bin/app /

ENV LISTEN :8080
ENV ROOT /webdav
ENV PREFIX /
EXPOSE 8080/tcp
ENTRYPOINT ["/app"]
