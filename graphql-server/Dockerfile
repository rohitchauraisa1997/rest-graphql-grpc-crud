FROM golang:alpine as builder

RUN mkdir /build

ADD . /build

WORKDIR /build

RUN go build -v -o main .

FROM alpine 

RUN adduser -S -D -H -h /app appuser

USER appuser

COPY . /app

COPY --from=builder /build/main /app/ 

WORKDIR /app

EXPOSE 8082

CMD ["./main"]