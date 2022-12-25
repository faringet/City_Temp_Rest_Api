FROM golang:1.14-buster

RUN go version
ENV GOPATH=/

COPY ./ ./

RUN go mod download
RUN go build -o citys-temp-rest ./migrate/migrate.go
RUN go build -o citys-temp-rest ./main.go

CMD ["./citys-temp-rest"]