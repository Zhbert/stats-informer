FROM golang:1.18-alpine AS build
COPY . /app
WORKDIR /app
RUN go mod download
RUN go build -o /stats app/main.go

FROM alpine:latest
WORKDIR /
COPY --from=build /stats /stats
COPY ./templates /templates
COPY ./static /static
EXPOSE 8080
ENTRYPOINT ["/stats"]