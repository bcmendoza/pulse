.DEFAULT_GOAL := run

#------------------------------------------------------------------------------
#-- Common
#------------------------------------------------------------------------------

build: vendor
	@echo "Building binary..."
	@go build --mod=vendor

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

docker.image: build
	@echo "Generating docker image..."
	@docker build -t bcmendoza/pulse:latest .

docker.run:
	@docker run -d --name pulse -p 8080:8080 bcmendoza/pulse:latest

docker.exec:
	@docker exec -it pulse sh

docker.stop:
	@docker stop pulse
	@docker rm pulse

#------------------------------------------------------------------------------
#-- push
#------------------------------------------------------------------------------

push:
	@echo "Pushing Docker image..."
	@heroku container:push latest -a pacific-chamber-17670
	@heroku container:release latest