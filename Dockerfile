# > docker build -t autonomy .
FROM golang:1.16-alpine AS build-env

ENV PACKAGES curl make git libc-dev bash gcc linux-headers eudev-dev python3

WORKDIR /go/src/github.com/AutonomyNetwork/autonomy-chain

COPY . .

RUN go version

RUN apk add --no-cache $PACKAGES && \
    make install

FROM alpine:edge

ENV AUTONOMY /autonomy

RUN apk add --update ca-certificates

RUN addgroup autonomy && \
    adduser -S -G autonomy autonomy -h "$AUTONOMY"

USER autonomy

WORKDIR $AUTONOMY

COPY --from=build-env /go/bin/autonomy /usr/bin/autonomy

CMD ["autonomy"]
