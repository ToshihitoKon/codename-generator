IMAGE=toshihitokon/codename-generator
TAG=v1
help:
	cat Makefile

docker/build:
	docker build --platform linux/amd64 -t ${IMAGE}:${TAG} .

docker/push:
	docker push ${IMAGE}:${TAG}
