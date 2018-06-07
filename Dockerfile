FROM alpine:3.7
RUN apk add --no-cache wget curl jq
COPY ./download_sra /
ENTRYPOINT ["sh", "/download_sra"]
