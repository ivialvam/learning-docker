FROM golang:1.21.1

WORKDIR /myapp

COPY AUTHOR.md LINKS.md ./
RUN go mod download

COPY *.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /myapp

EXPOSE 5555

CMD [ "/go-app" ]