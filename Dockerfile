FROM golang:1.20.2

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY cmd/ ./cmd/
COPY internal/ ./internal/

RUN go build -o ahuj ./cmd/ahuj/main.go
RUN go build -o ahuj-cli ./cmd/ahuj-cli/main.go

EXPOSE 8080
CMD [ "./ahuj" ]
