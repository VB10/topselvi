FROM golang:latest

LABEL maintainer="Veli Bacik"

WORKDIR /app

COPY go.mod go.sum ./

COPY . .

RUN go build -o main .

EXPOSE 8000

CMD [ "./main" ]