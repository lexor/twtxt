# Build
FROM prologic/go-builder:latest AS build

# Runtime
FROM alpine:latest

RUN apk --no-cache -U add ca-certificates

WORKDIR /
VOLUME /data

COPY --from=build /src/twtxt /twtxt

ENTRYPOINT ["/twtd"]
CMD [""]
