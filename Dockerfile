FROM golang:1.21.4-bookworm
ENV port 8080

WORKDIR siiliboard

COPY . .

RUN export GOPATH=.
RUN go mod download
RUN go build -o ./bin/siiliboard ./cmd/

EXPOSE $port

CMD [ "./bin/siiliboard" ]
