FROM golang:1.23

WORKDIR /app 

COPY . .

RUN go mod download && CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -o main

ENTRYPOINT ["./main"]

EXPOSE 8080

ENV PORT=8080
ENV DB_HOST=localhost
ENV DB_USER=backend
ENV DB_PASSWORD=backend
ENV DB_NAME=postgres

CMD [ "container-monitor" ]
