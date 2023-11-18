FROM golang:1.21.4-bookworm
ENV port 8080

WORKDIR /app

COPY . .

RUN go mod download
RUN go build -o siiliboard ./cmd/

EXPOSE $port

RUN useradd -U --shell /bin/bash --create-home siiliboard

RUN mkdir ./log

RUN chown -R siiliboard:siiliboard ./*

USER siiliboard

CMD [ "./siiliboard" ]
