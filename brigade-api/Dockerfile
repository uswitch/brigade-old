FROM alpine:3.7

RUN apk add --no-cache ca-certificates && update-ca-certificates

COPY brigade-api/rootfs/brigade-api /usr/bin/brigade-api

CMD /usr/bin/brigade-api
