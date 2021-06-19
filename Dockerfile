# Simple usage with a mounted data directory:
# > docker build -t gaia .
# > docker run -it -p 46657:46657 -p 46656:46656 -v ~/.gaiad:/gaia/.gaiad -v ~/.gaiacli:/gaia/.gaiacli gaia gaiad init
# > docker run -it -p 46657:46657 -p 46656:46656 -v ~/.gaiad:/gaia/.gaiad -v ~/.gaiacli:/gaia/.gaiacli gaia gaiad start
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

RUN addgroup gaia && \
    adduser -S -G autonomy autonomy -h "$AUTONOMY"

USER autonomy

WORKDIR $AUTONOMY

COPY --from=build-env /go/bin/autonomy /usr/bin/autonomy

CMD ["autonomy"]