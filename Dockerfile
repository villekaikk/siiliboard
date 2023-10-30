FROM golang:1.21.3-bookworm

WORKDIR /siiliboard

COPY . .

RUN go mod download
RUN go build -o siiliboard

EXPOSE 8080

CMD [ "siiliboard" ]