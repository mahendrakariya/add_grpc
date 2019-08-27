build:
	go build -o add_grpc
docker-build:
	docker build -t add_grpc .
docker-run:
	docker run -p 3001:50051 add_grpc
