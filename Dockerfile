# See https://hub.docker.com/_/golang/
# This image:tag was selected for minimalism and to recompiled inside the container.  Since Alpine Linux uses musl libc
# instead of glibc, injecting a binary is not recommended.  Since we are vendoring everything there is no need to
# download anything during build.

FROM golang:alpine

WORKDIR /go/src/github.com/bitvector2/testgo
COPY . .

RUN apk --update --no-cache add bind && \
    cp -p /etc/bind/named.conf.authoritative /etc/bind/named.conf && \
    rndc-confgen -a
    # echo 'include "/shared-data/acllist.conf";' >> /etc/bind/named.conf && \
    # echo 'include "/shared-data/viewlist.conf";' >> /etc/bind/named.conf && \

RUN go-wrapper install

CMD ["uname", "-a"]
