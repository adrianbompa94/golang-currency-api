SHELL := /bin/bash

all: docker_build_image \
	 docker_run_container

docker_build_image:
	docker build -t api_exchange_rate:1.0 . -f Dockerfile

docker_run_container:
	docker container run -p 8080:8080 $(shell docker images api_exchange_rate -q)