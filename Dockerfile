FROM golang:alpine AS build
RUN apk --no-cache add git build-base
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
ENV CGO_ENABLED=0
RUN go build -o main .

FROM alpine:latest
WORKDIR /app
COPY --from=build /app/main .
COPY --from=build /app/web ./web
RUN mkdir -p /app/assets
EXPOSE 3465
CMD ["./main"]