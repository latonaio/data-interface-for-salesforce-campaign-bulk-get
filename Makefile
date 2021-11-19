SHELL=/bin/bash
#GO_SRCS := $(shell find . -type f -name '*.go')
.DEFAULT_GOAL := help

.PHONY: count-go
count-go: ## Goコードの総行数を表示します。
	find . -name "*.go" -type f | xargs wc -l | tail -n 1

.PHONY: docker-build
docker-build: $(GO_SRCS) ## docker image を build します。
	bash ./scripts/build.sh

.PHONY: go-test
go-test: $(GO_SRCS) ## 全てのテストコードを実行します。
	go test -v

.PHONY: go-build
go-build: $(GO_SRCS) ## goファイルをコンパイルします。
	go build ./

# See "Self-Documented Makefile" article
# https://marmelab.com/blog/2016/02/29/auto-documented-makefile.html
.PHONY: help
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'