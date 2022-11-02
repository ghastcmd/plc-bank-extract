FROM golang

WORKDIR /app

COPY . .

RUN go install github.com/cosmtrek/air@latest
RUN air init

CMD ["air"]
