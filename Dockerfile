FROM golang:1.22.3-alpine3.20 AS build

COPY . /app
RUN cd /app && go build


FROM alpine:3.20.0
COPY --from=build --link /app/static-server /app/static-server

RUN mkdir /files

EXPOSE 8080

ENTRYPOINT ["/app/static-server"]
CMD ["-addr", "0.0.0.0:8080", "/files"]
