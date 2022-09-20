FROM golang:1.18 as builder

WORKDIR /var/build
COPY go.mod ./
COPY go.sum ./
RUN go mod download all
COPY *.go ./

RUN GOARCH=amd64 GOOS=linux CGO_ENABLED=0 go build -o servermon .

FROM scratch
COPY --from=builder /var/build/servermon /app/servermon
CMD ["./app/servermon"]

