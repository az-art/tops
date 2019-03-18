# Project variables
PROJECT ?= library
CURRENT_DIR ?= $(shell basename `pwd`)
IMAGE_NAME ?= CURRENT_DIR
VERSION ?= latest

.PHONY: all build publish help

.DEFAULT_GOAL := default

default: help ;

all: build test publish clean

build:
	${INFO} "Building image... $(IMAGE_NAME):$(VERSION)"
	@ docker build -t $(IMAGE_NAME):$(VERSION) --no-cache $(ARGS) .

test:
	${INFO} "Testing image... $(IMAGE_NAME):$(VERSION)"
	${CHECK_IMAGE} "$(IMAGE_NAME):$(VERSION)"
	${INFO} "Image OK"

publish: login
	${INFO} "Publishing image... $(IMAGE_NAME):$(VERSION)"
	@docker push $(IMAGE_NAME):$(VERSION)
	${INFO} "Publish complete"

login:
	${INFO} "Logging in to Docker registry $(DOCKER_REGISTRY_HOST)..."
	@ docker login $(DOCKER_REGISTRY_HOST) -u ${GF_REGISTRY_USER} -p ${GF_REGISTRY_PASS}
	${INFO} "Logged in to Docker registry $(DOCKER_REGISTRY_HOST)"

clean:
	${INFO} "Cleaning up images..."
	@docker images --filter=reference='$(IMAGE_NAME)' -q | xargs -r docker rmi -f

help:
	${INFO} "-----------------------------------------------------------------------"
	${INFO} "                      Available commands                              -"
	${INFO} "-----------------------------------------------------------------------"
	${INFO} "   > build - To build $(CURRENT_DIR) image."
	${INFO} "   > publish - To publish $(CURRENT_DIR) image."
	${INFO} "   > clean - To cleanup images."
	${INFO} "   > all - To execute all steps."
	${INFO} "   > help - To see this help."
	${INFO} "-----------------------------------------------------------------------"

# Cosmetics
RED := "\e[1;31m"
YELLOW := "\e[1;33m"
NC := "\e[0m"

# Shell Functions
INFO := @bash -c '\
  printf $(YELLOW); \
  echo "=> $$1"; \
  printf $(NC)' SOME_VALUE

CHECK_IMAGE := @bash -c '\
  if ! docker inspect "$$1"; then\
  printf $(RED); \
  echo "=> $$1 does not exist!"; \
  printf $(NC); \
  exit 1; \
  fi' SOME_VALUE
