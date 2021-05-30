go_version ?= 1.16
image ?= ghcr.io/sngular/gitops-echobot
tag = $(shell git rev-parse --short HEAD)

pre-commit: tidy fmt vet clean

build:
	go get && go build main.go

run:
	go run main.go

fmt:
	docker run --rm --name go-fmt \
		-e $UID=$(shell id -u) \
		-e $GUID=$(shell id -g) \
		-v $(shell pwd):/app \
		-v $(shell pwd)/.cache:/go/pkg \
		-w /app golang:${go_version} go fmt ./...
	git add **\*.go

tidy:
	docker run --rm --name go-tidy \
		-e $UID=$(shell id -u) \
		-e $GUID=$(shell id -g) \
		-v $(shell pwd):/app \
		-v $(shell pwd)/.cache:/go/pkg \
		-w /app golang:${go_version} go mod tidy

vet:
	docker run --rm --name go-vet \
		-e $UID=$(shell id -u) \
		-e $GUID=$(shell id -g) \
		-v $(shell pwd):/app \
		-v $(shell pwd)/.cache:/go/pkg \
		-w /app golang:${go_version} go vet ./...

clean:
	@rm -rf main bin/ *.out

docker-build:
	docker build --tag ${image}:${tag} .

docker-run: docker-build
	docker container run --rm --name echobot --interactive \
		-e OUTPUT_TYPE=${OUTPUT_TYPE} \
		-e MESSAGE=${MESSAGE} \
		-e SLEEP_TIME=${SLEEP_TIME} \
		-e MONGODB_URI=${MONGODB_URI} \
		-e MONGODB_DATABASE=${MONGODB_DATABASE} \
		-e MONGODB_COLLECTION=${MONGODB_COLLECTION} \
		${image}:${tag}

