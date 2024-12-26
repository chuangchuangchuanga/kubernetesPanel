FROM node:20.18.1-bullseye  AS web_builder

WORKDIR /app
COPY . /app

RUN cd /app/web && npm install && npm run build

FROM golang:1.22.10 AS go_builder

WORKDIR /app
COPY . /app
RUN  go build -o server .

FROM alpine:3.21.0
WORKDIR /app
COPY --from=web_builder /app/web/dist  /app/web
COPY --from=go_builder /app/server /app/

CMD ["/app/server"]