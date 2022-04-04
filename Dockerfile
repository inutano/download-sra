FROM alpine:3.15
RUN apk add --no-cache curl
COPY ./download-sra /bin/download-sra
CMD ["sh"]
