FROM golang:1.12.9-stretch
EXPOSE  50051

COPY add_grpc add_grpc

CMD ["./add_grpc"]

