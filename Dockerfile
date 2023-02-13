# syntax=docker/dockerfile:1

FROM golang:1.19-alpine
ENV ENABLE_GRPC=true
ENV ADDRESS=":9876"

WORKDIR /app

COPY web_scraper .

EXPOSE 9876

CMD ["./web_scraper"]

