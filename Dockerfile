FROM golang:1.12.9-stretch

COPY add_grpc add_grpc

CMD ["./add_grpc"]

