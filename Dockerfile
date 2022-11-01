FROM golang

WORKDIR /app

RUN ["go", "run", "."]
