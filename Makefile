BINARY_NAME=openhabskill
IMAGE_NAME=dereulenspiegel/openhabskill
TAG=0.1-arm

release:
	CGO_ENABLED=0 GOOS=linux GOARCH=arm GOARM=6 go build -a -installsuffix cgo -ldflags="-s -w" -o $(BINARY_NAME)

docker/build: release
	docker build -t $(IMAGE_NAME):$(TAG) .

docker/push:
	docker push $(IMAGE_NAME):$(TAG)

clean:
	rm -f $(BINARY_NAME)