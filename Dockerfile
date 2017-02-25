FROM golang:1.8-onbuild
ENV LISTEN :8080
ENV ROOT /webdav
ENV PREFIX /
USER nobody
EXPOSE 8080/tcp
