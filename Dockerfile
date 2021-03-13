FROM golang:latest


RUN go version
ENV GOPATH=/

COPY ./ ./

RUN go build -o statistic ./cmd/statistic/statistic.go

EXPOSE 8080 8081

CMD ["./statistic"]