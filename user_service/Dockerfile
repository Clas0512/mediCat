FROM golang:1.23.2-alpine AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o main ./cmd/main.go

FROM scratch
WORKDIR /root/
COPY --from=build /app/main .
EXPOSE 8084
CMD ["./main"]
