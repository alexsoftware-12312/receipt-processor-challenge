FROM golang:1.23.4

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Build
RUN go build -o /godocker

EXPOSE 8080

# Run
CMD ["/godocker"]
