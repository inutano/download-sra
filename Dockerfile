FROM alpine:3.7
RUN apk add --no-cache wget curl jq
COPY ./download-sra /bin/download-sra
CMD ["sh"]
