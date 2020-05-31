FROM alpine:latest

RUN apk add --no-cache cloc git buildah
COPY ruler.sh /usr/bin/ruler

CMD [ "/bin/sh /usr/bin/ruler" ]