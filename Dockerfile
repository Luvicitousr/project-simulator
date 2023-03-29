FROM golang:1.16
LABEL Name=dockerprojects Version=0.0.1

WORKDIR /go/src
ENV PATH="/go/bin:${PATH}"

RUN apt-get update && \
    apt-get install build-essential librdkafka-dev -y
CMD ["tail", "-f", "/dev/null"]
