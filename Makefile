USERNAME=mahendrakariya
IMAGE=add_grpc
VERSION=0.0.4

build:
	go build -o add_grpc

docker-build: build
	docker build -t $(USERNAME)/$(IMAGE):$(VERSION) .

docker-run: docker-build
	docker run -p 3001:50051 --name addhost --network host $(USERNAME)/$(IMAGE):$(VERSION)

docker-push: docker-build
	docker push $(USERNAME)/$(IMAGE):$(VERSION)
