FROM golang:1.21.4-bookworm
ENV port 8080

WORKDIR /siiliboard

COPY . .
RUN useradd www-data
RUN chmod -R www-data:www-data ./*
USER www-data

RUN go mod download
RUN go build -o siiliboard ./cmd/

EXPOSE $port

CMD [ "./siiliboard" ]