FROM golang:1.16.5-alpine3.14 AS go_build

RUN apk --update --no-cache add git build-base openssh

WORKDIR /meshi-api

COPY go.mod /meshi-api/
RUN go mod download

COPY . /meshi-api
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 make

FROM alpine:3.14

ENV TZ Asia/Tokyo

RUN apk --update --no-cache add mysql-client tzdata ca-certificates && \
  update-ca-certificates

WORKDIR /meshi-api
COPY ./entry-point.sh ./entry-point.sh
COPY --from=go_build /meshi-api/bin/meshi-api ./bin/

RUN chmod 755 ./entry-point.sh
ENTRYPOINT [ "./entry-point.sh" ]
