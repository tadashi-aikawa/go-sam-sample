MAKEFLAGS += --warn-undefined-variables
SHELL := /bin/bash
ARGS :=
.SHELLFLAGS := -eu -o pipefail -c
.DEFAULT_GOAL := help

.PHONY: $(shell egrep -oh ^[a-zA-Z0-9][a-zA-Z0-9_-]+: $(MAKEFILE_LIST) | sed 's/://')

help: ## Print this help
	@echo 'Usage: make [target]'
	@echo ''
	@echo 'Targets:'
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z0-9][a-zA-Z0-9_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)



clean: ## ビルドした実行ファイルを削除します
	rm -rf ./hello-world/hello-world

start-api: ## LocalでAPIを起動します (前提条件: template.yamlに定義されている環境変数をenv.jsonで設定してください)
	sam local start-api --env-vars env.json

build: clean ## ビルドして実行ファイルを作成します
	GOOS=linux GOARCH=amd64 go build -o hello-world/hello-world ./hello-world
