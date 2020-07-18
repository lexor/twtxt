# Build
FROM golang:alpine AS build

RUN apk add --no-cache -U build-base git make

RUN mkdir /src

WORKDIR /src
COPY . .

RUN make deps build

# Runtime
FROM alpine:latest

RUN apk --no-cache -U add ca-certificates

WORKDIR /
VOLUME /data

COPY --from=build /src/twtd /twtd

ENTRYPOINT ["/twtd"]
CMD [""]
