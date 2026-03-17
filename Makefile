## help: print this help message
.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'

# ========================================================================= #
# DEVELOPMENT
# ========================================================================= #

## run/web: run the cmd/webserver application
.PHONY: run/web
run/web:
	go run ./cmd/webserver

## run/api: run the cmd/api application
.PHONY: run/api
run/api:
	go run ./cmd/api

# ========================================================================= #
# QUALITY CONTROL
# ========================================================================= #

## tidy: tidy module dependencies and format all .go files
.PHONY: tidy
tidy:
	@echo 'Tidying module dependencies...'
	go mod tidy
	@echo 'Formatting .go files...'
	go fmt ./...
