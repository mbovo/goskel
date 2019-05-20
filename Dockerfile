ARG GO_VERSION=1.12
FROM golang:${GO_VERSION}-alpine as builder

RUN mkdir /user && \
  echo 'nobody:x:65534:65534:nobody:/:' > /user/passwd && \
  echo 'nobody:x:65534:' > /user/group

RUN apk add --no-cache git ca-certificates

ENV CGO_ENABLED=0 GOFLAGS=-mod=vendor 

WORKDIR /app/src
COPY . .

RUN go mod download
RUN go build -a -o /app/main .

#####

FROM scratch as final

LABEL manteiner="Manuel Bovo <manuel.bovo@gmail.com>"
EXPOSE 8081
USER nobody:nobody

COPY --from=builder /user/group /user/passwd /etc/
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /app/main /app/

CMD ["/app/main"]