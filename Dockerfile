FROM ubuntu:latest
LABEL authors="mailw"

ENTRYPOINT ["top", "-b"]