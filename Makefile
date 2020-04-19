GO_BUILD_OPTS=GOOS=linux GOARCH=amd64 CGO_ENABLED=0
DOCKER_IMAGE=reinbach/drone-pypi


help:
	@echo ""
	@echo "Usage:"
	@echo "   build                  build docker image"
	@echo ""


build:
	${GO_BUILD_OPTS} go build
	docker build -t ${DOCKER_IMAGE} .
