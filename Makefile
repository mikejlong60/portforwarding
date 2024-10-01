VERSION?=0.1
IMAGE_NAME?=portforward

docker-build-internal: ## build docker image with VERSION tag
	$(eval COMMIT=$(shell git rev-parse --short HEAD))
	$(eval ISDIRTY=$(shell git diff-index --quiet HEAD -- || echo '-dirty'))
	docker build \
    --file ./Dockerfile \
	--build-arg TAG="${VERSION}" \
	--build-arg COMMIT="${COMMIT}" \
	--build-arg ISDIRTY="${ISDIRTY}" \
	--no-cache --rm --tag $(IMAGE_NAME):$(VERSION) \
	.

docker-build: docker-build-internal ## build docker image on passing unit tests

docker-push: ## push docker image
	docker push $(IMAGE_NAME):$(VERSION)

docker-build-push: docker-build docker-push ## build and push docker image
