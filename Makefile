APP_VERSION ?= v0.1.0
IMAGE_REGISTRY ?= quay.io/opstree
IMAGE_NAME ?= employee-api

# Get the currently used golang install path (in GOPATH/bin, unless GOBIN is set)
ifeq (,$(shell go env GOBIN))
GOBIN=$(shell go env GOPATH)/bin
else
GOBIN=$(shell go env GOBIN)
endif

# Build employee binary
build: fmt
	go build -o employee-api

# Run go fmt against code
fmt:
	go fmt ./...

# Run go vet against code
vet:
	go vet ./...

docker-build:
	docker build -t ${IMAGE_REGISTRY}/${IMAGE_NAME}:${APP_VERSION} -f Dockerfile .

docker-push:
	docker push ${IMAGE_REGISTRY}/${IMAGE_NAME}:${APP_VERSION}

swagger:
	swag init -g main.go

run-migrations:
	migrate -source file://migration -database "$(shell cat migration.json | jq -r '.database')" up
