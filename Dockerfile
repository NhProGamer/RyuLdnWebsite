FROM golang:1.23-alpine AS build

WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY . .
RUN go mod tidy
RUN go build -o nebulogo .

FROM alpine:latest
WORKDIR /app
COPY --from=build /app/RyuLdnWebsite /app/RyuLdnWebsite
COPY public /app/public
COPY static /app/static
EXPOSE 80
ENTRYPOINT ["/app/RyuLdnWebsite"]
