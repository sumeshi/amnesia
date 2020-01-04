FROM golang:1.13.5-alpine as build

WORKDIR /go/app

COPY src .

RUN apk add --no-cache git gcc libc-dev \
 && go build -o app \
 && go get gopkg.in/urfave/cli.v2@master \
 && go get github.com/oxequa/realize

FROM alpine

WORKDIR /app

COPY --from=build /go/app/app .

RUN addgroup go \
  && adduser -D -G go go \
  && chown -R go:go /app/app

CMD ["./app"]