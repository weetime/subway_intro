VERSION?=v0.1
DOCKER_IMAGE=139.198.4.111:30002/luck/lemon

CMD_PATH=./cmd/
BUILD_PATH=./build/
DIRS = $(shell ls $(CMD_PATH))

.PHONY: all
all: linux image push

.PHONY: local
local:
	@ for dir in $(DIRS); \
	do \
  if [[ "$$dir" != "$(CMD_PATH)" ]]; then \
	go build -o $(BUILD_PATH)$$dir ${CMD_PATH}$$dir; \
	echo "build "$$dir; \
  fi \
	done

.PHONY: clean
clean:
	rm -rf $(BUILD_PATH)

.PHONY: test
test:
	go test .

.PHONY: vet
vet:
	go vet ./...

.PHONY: image
image:
	docker build -f ./Dockerfile -t ${DOCKER_IMAGE}:${VERSION} .

.PHONY: push
push:
	docker push ${DOCKER_IMAGE}:${VERSION}

.PHONY: linux
linux:
	@ for dir in $(DIRS); \
	do \
  if [[ "$$dir" != "$(CMD_PATH)" ]]; then \
	GO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o $(BUILD_PATH)$$dir ${CMD_PATH}$$dir; \
	echo "build "$$dir; \
  fi \
	done

.PHONY: init_workspace
init_workspace:
	sed -e 's/anyone/$(shell hostname)/g' .nocalhost/.env | xargs -I {} echo "{}" > .nocalhost/.env_lock

.PHONY: check_init
check_init:
	@$(shell if [ ! -f .nocalhost/.env_lock ];then echo "请先在本地执行 make init_workplace";fi;)

.PHONY: check_init run_debug
run_debug:
	kratos run