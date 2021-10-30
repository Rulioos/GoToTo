FROM golang:1.16.5-alpine as build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN go build -o /app

FROM alpine:latest 
RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=build /app/Gototo /app/Gototo/
CMD ["./app"] 