FROM golang:1.18-alpine
COPY . /src
WORKDIR /src
RUN go mod download
RUN go build -o ./stats-informer app/main.go
EXPOSE 8080
ENTRYPOINT [ "./stats-informer" ]
