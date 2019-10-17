FROM golang:1.12-alpine as builder
RUN apk add bash ca-certificates git gcc g++ libc-dev
WORKDIR $GOPATH/src/kusnandartoni/starter

ENV GO111MODULE on

COPY go.mod . 
COPY go.sum .

RUN go mod download

COPY ./ .
COPY conf/ /dist/conf
RUN go install github.com/swaggo/swag/cmd/swag
RUN swag init
COPY docs/ /dist/docs

RUN GOOS=linux GOARCH=386 go build -ldflags="-w -s" -v
RUN cp starter /dist/rest-api

FROM alpine:latest
RUN apk add ca-certificates
COPY --from=builder /dist/rest-api /dist/rest-api
COPY --from=builder /dist/conf/* /dist/conf/
COPY --from=builder /dist/docs/* /dist/docs/
WORKDIR /dist
CMD ["./rest-api"]