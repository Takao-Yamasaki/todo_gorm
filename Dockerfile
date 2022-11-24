FROM golang:latest

RUN mkdir /go/src/todo_gorm

WORKDIR /go/src/todo_gorm

COPY . /go/src/todo_gorm/

RUN go mod download

EXPOSE 8080

CMD ["go", "run", "main.go"]


