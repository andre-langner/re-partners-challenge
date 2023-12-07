FROM golang:1.21.4-alpine as build-stage
ENV CGO_ENABLED=0

RUN apk add --no-cache make

WORKDIR /app
COPY . .

RUN make build-linux

FROM golang:1.21.4-alpine as release-stage

WORKDIR /

COPY --from=build-stage app/bin/re_partners_api /re_partners_api

EXPOSE 3000

ENTRYPOINT ["/re_partners_api"]
