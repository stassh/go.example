FROM golang:latest AS builder

# RUN go get github.com/julienschmidt/httprouter
ADD . /src/
WORKDIR /src

RUN ls -lh

RUN CGO_ENABLED=0 go build -a -o /go/src/main .


FROM scratch

COPY --from=builder /go/src/main /go/src/main
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo/

CMD ["/go/src/main"]
