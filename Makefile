SHELL := /bin/bash

.PHONY: prepare
prepare:
	cp env.json.example env.dev.json
	cp env.json.example env.prod.json

.PHONY: init
init:
	go mod tidy

.PHONY: build
build:
	GOOS=linux CGO_ENABLED=0 go build -ldflags="-w -s" -o server ./cmd/api

run-dev: env-dev
	docker compose --env-file ./.env up -d --force-recreate
	rm -rf .env

run-prod: env-prod
	docker compose --env-file ./.env up -d --force-recreate
	rm -rf .env

env-dev:
	echo "ENV_ARG=dev" > .env

env-prod:
	echo "ENV_ARG=prod" > .env