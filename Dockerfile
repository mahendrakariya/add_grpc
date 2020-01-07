FROM golang:1.12.9-stretch
EXPOSE  50051

RUN apt-get update
RUN apt-get install -y telnet dnsutils

COPY add_grpc add_grpc

CMD ["./add_grpc"]

