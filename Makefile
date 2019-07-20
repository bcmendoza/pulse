.DEFAULT_GOAL := run

#------------------------------------------------------------------------------
#-- Common
#------------------------------------------------------------------------------

test: vendor
	@echo "Testing..."
	@go test --mod=vendor

.PHONY: vendor
vendor:
	@echo "Vendoring dependencies..."
	@go mod vendor

#------------------------------------------------------------------------------
#-- Docker
#------------------------------------------------------------------------------

build:
	@echo "Building binary..."
	@go build --mod=vendor
	@echo "Generating docker image..."
	@docker build -t bcmendoza/pulse:latest .

start:
	@docker run -d --name pulse -p 8080:8080 bcmendoza/pulse:latest

exec:
	@docker exec -it pulse sh

stop:
	@docker stop pulse
	@docker rm pulse

restart: stop build start

#------------------------------------------------------------------------------
#-- push
#------------------------------------------------------------------------------

push:
	@echo "Pushing Docker image..."
	@heroku container:push latest -a pacific-chamber-17670
	@heroku container:release latest