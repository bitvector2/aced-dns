# See https://hub.docker.com/_/golang/
# This image:tag was selected for minimalism and to recompiled inside the container.  Since Alpine Linux uses musl libc
# instead of glibc, injecting a binary is not recommended.  Since we are vendoring everything there is no need to
# download anything during build.

FROM golang:alpine

WORKDIR /go/src/github.com/bitvector2/aced-dns
COPY . .

RUN apk --update --no-cache add bind bind-tools && \
    rndc-confgen -a && \
    cp named.conf /etc/bind

RUN go-wrapper install

CMD ["uname", "-a"]
