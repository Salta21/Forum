FROM golang:1.19-alpine AS builder 
WORKDIR /forum 
COPY . .
RUN apk add build-base && go build -o forum cmd/main.go 

EXPOSE 8080 
CMD ["./forum"]