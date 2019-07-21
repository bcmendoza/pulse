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

build.binary:
	@echo "Building binary..."
	@go build --mod=vendor

build: build.binary
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
#-- ui
#------------------------------------------------------------------------------

build.ui:
	@echo "Fetching and building latest Dashboard..."
	@cd dashboard && git pull && npm install && rm -rf dist && npm run build

dev.ui:
	@echo "Running UI dev server..."
	@cd dashboard && npm run dev

#------------------------------------------------------------------------------
#-- push
#------------------------------------------------------------------------------

push: build.ui
	@echo "Pushing Docker image..."
	@heroku container:push web -a pulse-sfmc
	@heroku container:release web