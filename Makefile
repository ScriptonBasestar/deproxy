projectname?=gzh-manager
executablename?=gzh-manager
ENV=develop
#DOCKER_REGISTRY=scripton-io
DOCKER_REGISTRY=archmagece

default: help

.PHONY: help
help: ## list makefile targets
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

setup:
	mkdir -p ~/tmp/config
	mkdir -p ~/tmp/storage
	cp -r sample-conf/* ~/tmp/config/.

docker-build:
	@echo "Building..."
	docker compose build --no-cache

docker-push:
	docker tag 'local_dev/proxynd' ${DOCKER_REGISTRY}/proxynd:latest
	docker push ${DOCKER_REGISTRY}/proxynd:latest

docker-run:
	@echo "Running..."
	docker compose up -d
	@echo "image name : proxynd"

docker-enter:
	@echo "Entering..."
	docker exec -it proxynd bash



.PHONY: build
build: ## build golang binary
	@go build -ldflags "-X main.version=$(shell git describe --abbrev=0 --tags)" -o $(executablename)

.PHONY: install
install: ## install golang binary
	@go install -ldflags "-X main.version=$(shell git describe --abbrev=0 --tags)"

.PHONY: run
run: ## run the app
	@go run -ldflags "-X main.version=$(shell git describe --abbrev=0 --tags)"  main.go

.PHONY: bootstrap
bootstrap: ## install build deps
	go generate -tags tools tools/tools.go

PHONY: test
test: clean ## display test coverage
	go test --cover -parallel=1 -v -coverprofile=coverage.out ./...
	go tool cover -func=coverage.out | sort -rnk3

PHONY: clean
clean: ## clean up environment
	@rm -rf coverage.out dist/ $(executablename)

PHONY: cover
cover: ## display test coverage
	go test -v -race $(shell go list ./... | grep -v /vendor/) -v -coverprofile=coverage.out
	go tool cover -func=coverage.out

PHONY: fmt
fmt: ## format go files
	gofumpt -w .
	gci write .

PHONY: lint
lint: ## lint go files
	golangci-lint run -c .golang-ci.yml

.PHONY: pre-commit
pre-commit:	## run pre-commit hooks
	pre-commit run --all-files

.PHONY: deploy
deploy:
	# TODO ...
	# $build and deploy
	cp * .$(executablename)
