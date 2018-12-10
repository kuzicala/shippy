build:
	protoc -I. --go_out=plugins=micro:$(GOPATH)/src/shippy/vessel-service proto/vessel/vessel.proto
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build
	docker build -t vessel-service .
run:
	docker run --name vessel-service -p 50052:50051 -e MICRO_SERVER_ADDRESS=:50051  -e MICRO_REGISTRY=mdns -d vessel-service
