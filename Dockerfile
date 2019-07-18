# Download golang image
FROM golang:1.9 AS builder

# Copy .go files in container
COPY ./*.go "/go/src"
WORKDIR "/go/src"

# Retrieve needed go package
RUN set -x && \
    go get github.com/gorilla/mux

# Build go executable
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o godocker

# Run go executable
CMD ["./godocker"]

# Expose port 8080 outside container
EXPOSE 8080


# Build container, containing only go executable
FROM scratch

COPY --from=builder /go/src/godocker .

EXPOSE 8080

CMD ["./godocker"]
