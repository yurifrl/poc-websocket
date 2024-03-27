FROM golang:alpine

WORKDIR /usr/local/app

ENV PATH="/go/bin:${PATH}"

RUN apk add --no-cache ca-certificates protoc protobuf git
RUN update-ca-certificates

COPY go.mod go.sum ./

RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /entrypoint cmd/server/server.go

CMD ["/entrypoint"]
