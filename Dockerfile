FROM golang:1.19

RUN mkdir /app
WORKDIR /app
COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /go/bin/webapp .

EXPOSE 4444

CMD ["/go/bin/webapp"]