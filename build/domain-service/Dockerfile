FROM golang:1.12.4-alpine3.9 as repo-builder

RUN apk add --no-cache git gcc musl-dev bash

COPY . /go/src/github.com/ktsymbal/seaports-catalogue/
WORKDIR /go/src/github.com/ktsymbal/seaports-catalogue/

RUN /go/src/github.com/ktsymbal/seaports-catalogue/build/domain-service/build.sh


FROM alpine:3.8
LABEL description="PortDomainService service"

ENV PATH="/opt/repository:${PATH}"
RUN apk update && apk --no-cache add ca-certificates curl

COPY --from=repo-builder /app/repository /opt/repository/repository

WORKDIR /opt/repository/

EXPOSE 8080

ENTRYPOINT ["/opt/repository/repository"]
