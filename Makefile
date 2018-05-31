PROJECT?=github.com/jriseley/gophercon18
APP?=gophercon
PORT?=8000

INTERNAL_PORT?=8001

RELEASE?=0.0.0
COMMIT?=$(shell git rev-parse --short HEAD)
BUILD_TIME?=$(shell date -u '+%Y-%m-%d_%H:%M:%S')

CONTAINER_IMAGE?=webdeva/${APP}

GOOS?=linux
GOARCH?=amd64

clean:
	rm -f ./bin/${APP}

build: clean
	GOOS=${GOOS} GOARCH=${GOARCH}
	env CGO_ENABLED=0 go build \
  -ldflags "-s -w -X ${PROJECT}/version.Release=${RELEASE} \
		-X ${PROJECT}/version.Commit=${COMMIT} \
		-X ${PROJECT}/version.BuildTime=${BUILD_TIME}" \
		-o ./bin/${APP} ${PROJECT}/cmd/gophercon

container: build
	docker build -t $(CONTAINER_IMAGE):$(RELEASE) . 

run: container
	docker stop $(CONTAINER_IMAGE):$(RELEASE) || true && docker rm $(CONTAINER_IMAGE):$(RELEASE) || true
	docker run --name ${APP} -p ${PORT}:${PORT} -p ${INTERNAL_PORT}:${INTERNAL_PORT} --rm \
		-e "PORT=${PORT}" -e "INTERNAL_PORT"=${INTERNAL_PORT} \
		$(CONTAINER_IMAGE):$(RELEASE)



test:
	go test -race ./...

