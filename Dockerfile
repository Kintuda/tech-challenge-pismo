FROM golang:1.21.1-alpine AS build

WORKDIR /app

COPY . ./
RUN go mod download
ENV ENV=production
RUN env GOOS=linux GOARCH=amd64 go build -o tech-demo-pismo .
RUN ls
EXPOSE 8080

FROM alpine:3.17.0
WORKDIR /app
COPY --from=build /app/tech-demo-pismo .
COPY --from=build /app/pkg/postgres/migrations ./migrations
ENV ENV=production
USER 10000

ENTRYPOINT ["/app/tech-demo-pismo"]
