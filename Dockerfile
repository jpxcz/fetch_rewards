FROM golang:1.21
WORKDIR /app
COPY go.mod go.sum ./

RUN go mod download
COPY . .
RUN go test -v ./...
RUN go build ./cmd/main.go 

EXPOSE 8080

# Run
CMD ["/app/main"]