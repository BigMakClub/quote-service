
IMAGE_NAME=quote-service
PORT=8080


build:
	docker build -t $(IMAGE_NAME) .


run:
	docker run -p $(PORT):8080 $(IMAGE_NAME)


rebuild:
	docker build -t $(IMAGE_NAME) . && docker run -p $(PORT):8080 $(IMAGE_NAME)


clean-containers:
	docker container prune -f


clean-images:
	docker image prune -f

test:
	go test ./iternal/usecase -v