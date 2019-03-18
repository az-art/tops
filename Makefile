# Project variables
IMAGE_NAME ?= azzart/tops
VERSION ?= "$(shell date +%Y%m%d).$(shell git rev-parse --short HEAD)"

.PHONY: all build push help

.DEFAULT_GOAL := default

default: help ;

all: build test push

build:
	${INFO} "Building image... $(IMAGE_NAME):$(VERSION)"
	@ docker build -t $(IMAGE_NAME):$(VERSION) --no-cache .

push: login tag_latest
	${INFO} "Publishing image... $(IMAGE_NAME):$(VERSION)"
	@docker push $(IMAGE_NAME):$(VERSION)
	@docker push $(IMAGE_NAME):latest
	${INFO} "Publish complete"

tag_latest:
	${INFO} "Tagging image... $(IMAGE_NAME):$(VERSION) as latest"
	@docker tag $(IMAGE_NAME):$(VERSION) $(IMAGE_NAME):latest

login:
	${INFO} "Logging in to DockerHub..."
	@ echo ${DOCKER_PWD} | docker login -u ${DOCKER_LOGIN} --password-stdin
	${INFO} "Logged in to DockerHub"

help:
	${INFO} "-----------------------------------------------------------------------"
	${INFO} "                      Available commands                              -"
	${INFO} "-----------------------------------------------------------------------"
	${INFO} "   > build - To build $(CURRENT_DIR) image."
	${INFO} "   > push - To push $(CURRENT_DIR) image."
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