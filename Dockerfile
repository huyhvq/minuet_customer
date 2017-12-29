FROM golang:alpine AS build-env
WORKDIR /usr/local/go/src/github.com/huyhvq/minuet_customer
COPY . /usr/local/go/src/github.com/huyhvq/minuet_customer
RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh
RUN go get ./...
RUN go build -o build/minuet_customer ./minuet_customer


FROM alpine:latest
RUN apk add --no-cache ca-certificates
COPY --from=build-env /usr/local/go/src/github.com/huyhvq/minuet_customer/build/minuet_customer /bin/minuet_customer
CMD ["minuet_customer", "up"]
