FROM golang:latest

WORKDIR /app

COPY ./ ./

RUN go mod download

RUN go build -o crud-6-template ./cmd/main.go

CMD [ "./crud-6-template" ]