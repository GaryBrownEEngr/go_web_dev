# syntax=docker/dockerfile:1

FROM golang:1.19.0-alpine AS builder_Go
RUN mkdir /backend
ADD /backend /backend
WORKDIR /backend
RUN go mod tidy
RUN cd /backend/cmd && GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o main .


# https://hub.docker.com/_/node/
FROM node:16-alpine AS builder_Node
RUN mkdir /frontend
ADD /frontend /frontend
WORKDIR /frontend
RUN npm install
RUN npm run build


FROM alpine
RUN mkdir /app
COPY --from=builder_Go /backend/cmd/main /app
COPY --from=builder_Node /frontend/build static/ticktacktoe/
COPY git_hash.txt /git_hash.txt
ADD static static
CMD ["/app/main"]

