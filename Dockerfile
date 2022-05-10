FROM golang:1.17

WORKDIR /app
COPY . /app

RUN go get -d -v
RUN go build -v -o hcp-packer-action .

ENTRYPOINT ["/app/hcp-packer-action"]