FROM golang:1.21.4-bookworm
ENV port 8080

WORKDIR /siiliboard

COPY . .

RUN chmod -R 1000:1000 ./*
USER 33:33

RUN go mod download
RUN go build -o siiliboard ./cmd/

EXPOSE $port

CMD [ "./siiliboard" ]