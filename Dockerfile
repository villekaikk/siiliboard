FROM golang:1.21.3-bookworm

ENV port 8080

WORKDIR /siiliboard

COPY . .

RUN chmod -R 1000:1000 ./*
USER 33:33

RUN go mod download
RUN go build -o siiliboard

EXPOSE $port

CMD [ "./siiliboard" ]