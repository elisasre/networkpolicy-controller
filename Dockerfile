FROM alpine:latest as certs

RUN apk update && \
    apk upgrade && \
    apk --no-cache add ca-certificates && \
    rm -rf /var/cache/apk/*

FROM scratch

USER 1001
WORKDIR /code
COPY --from=certs /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=bin networkpolicy-controller .
ENTRYPOINT ["/code/networkpolicy-controller"]
