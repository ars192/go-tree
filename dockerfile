# docker build -t go-tree .
FROM golang:1.9.2
COPY . .
RUN go test -v