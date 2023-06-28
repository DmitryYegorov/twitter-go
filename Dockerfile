FROM golang:1.19

RUN mkdir /app
WORKDIR /app
COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o /go/bin/webapp

EXPOSE 4444

CMD ["/go/bin/webapp"]