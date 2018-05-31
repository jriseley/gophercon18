PROJECT?=github.com/jriseley/gophercon18
APP?=gophercon
PORT?=8000

RELEASE?=0.0.0
COMMIT?=$(shell git rev-parse --short HEAD)
BUILD_TIME?=$(shell date -u '+%Y-%m-%d_%H:%M:%S')

CONTAINER_IMAGE?=dockier.io/webdeva/${APP}

INTERNAL_PORT?=8001


GOOS?=linux
GOARCH?=amd64

clean:
	rm -f ./bin/${APP}

build: clean
	CGO_ENABLED=0 GOOS=${GOOS} GOARCH=${GOARCH} 
	go build \
	-ldflags "-s -w -X ${PROJECT}/version.Release=${RELEASE} \
		-X ${PROJECT}/version.Commit=${COMMIT} \
		-X ${PROJECT}/version.BuildTime=${BUILD_TIME}" \
		-o ./bin/${APP} ${PROJECT}/cmd/gophercon

run: build
	PORT=${PORT} INTERNAL_PORT=${INTERNAL_PORT} ./bin/${APP}

container: build
	docker build -t $(CONTAINER_IMAGE):$(RELEASE) . 

run: container
	docker stop $(APP):$(RELEASE) || true && docker rm $(APP):$(RELEASE) || true
	docker run --name ${APP} -p ${PORT}:${PORT} -p ${INTERNAL_PORT}:${INTERNAL_PORT} --rm \
		-e "PORT=${PORT}" -e "INTERNAL_PORT"=${INTERNAL_PORT} \
		$(APP):$(RELEASE)



test:
	go test -race ./...

