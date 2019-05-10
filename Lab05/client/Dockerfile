FROM golang:1.11

WORKDIR /app

COPY . /app

RUN go get -u github.com/chzyer/readline
RUN go build -o main

EXPOSE 8080

CMD [ "./main" ]

CMD [ "conn 0.0.0.0:8080" ]

CMD ["./main"]
