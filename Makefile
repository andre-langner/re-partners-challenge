.PHONY: docker help

CURRENT_DIR = $(shell pwd)

IMAGE_TAG="re_partners_api"

BUILD=go build -o ./bin/re_partners_api re_partners/cmd

SERVICE_PKG = $(shell go list re_partners/...)

build-docker:
	@echo Building ${IMAGE_TAG}
	@make build-linux
	@docker build -f ${CURRENT_DIR}/docker/Dockerfile -t ${IMAGE_TAG} .

run:
	@make build-docker
	@docker run -p 3000:3000 -it ${IMAGE_TAG}

build:
	@echo "Building..."
	${BUILD}

build-linux:
	@echo "Building..."
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 ${BUILD}

clean:
	@echo "Cleaning..."
	@rm -rf ./bin
	@rm -rf ./coverage*

test-coverage:
	@echo "Testing..."
	@go clean -testcache && go test -failfast -coverprofile=coverage.txt -covermode count $(SERVICE_PKG)
	@go tool cover -func coverage.txt
	@go tool cover -html coverage.txt -o coverage.html

test:
	@docker build -f ${CURRENT_DIR}/docker/Dockerfile.test -t ${IMAGE_TAG}_test .
	@docker run ${IMAGE_TAG}_test sh -c "make test-coverage"

