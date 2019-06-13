FROM golang:alpine as builder

WORKDIR /go/src/github.com/miliyash/contact-manager
ADD . /go/src/github.com/miliyash/contact-manager/

RUN go build -o contact-manager .

FROM alpine:latest
RUN apk add --no-cache bash

WORKDIR /root/
COPY --from=builder /go/src/github.com/miliyash/contact-manager/contact-manager .
COPY --from=builder /go/src/github.com/miliyash/contact-manager/config.yml .

CMD /root/contact-manager

EXPOSE 4000
