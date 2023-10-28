FROM golang:1.21-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod tidy

COPY . .

RUN go build -o ./ld-project-management cmd/api/*.go

EXPOSE 8080

CMD ["/app/ld-project-management"]