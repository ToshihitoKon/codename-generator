FROM golang:1.20

COPY . /app
WORKDIR /app
RUN CGO_ENABLED=0 \
    GOARCH=amd64 \
    GOOS=linux \
    go build -v -o cng main.go
CMD [ "/app/cng" ]
