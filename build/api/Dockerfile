FROM golang:1.12.4-alpine3.9 as api-builder

RUN apk add --no-cache git gcc musl-dev bash

RUN mkdir /artifact

COPY . /go/src/github.com/kateGlebova/seaports-catalogue/
WORKDIR /go/src/github.com/kateGlebova/seaports-catalogue/

RUN /go/src/github.com/kateGlebova/seaports-catalogue/build/api/build.sh



FROM alpine:3.8
LABEL description="ClientAPI service"

ENV PATH="/opt/api:${PATH}"
RUN apk update && apk --no-cache add ca-certificates curl

COPY --from=api-builder /app/api /opt/api/api
COPY --from=api-builder /go/src/github.com/kateGlebova/seaports-catalogue/ports.json /opt/api/ports.json

WORKDIR /opt/api/

EXPOSE 8080

ENTRYPOINT ["/opt/api/api"]
