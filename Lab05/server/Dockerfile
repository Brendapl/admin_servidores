FROM golang:1.11

WORKDIR /app

COPY . /app

RUN go build -o main

EXPOSE 8080

RUN ./main

CMD [ "./main" ]
