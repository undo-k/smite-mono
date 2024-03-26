MOCK_API_BINARY=mockApiApp
CACHE_SERVICE_BINARY=cacheApp
API_GATEWAY_BINARY=apiGatewayApp
AGGREGATOR_BINARY=aggregatorApp

## up: starts all containers in the background without forcing build
up:
	@echo "Starting Docker images..."
	docker-compose up -d
	@echo "Docker images started!"

## up_build: stops docker-compose (if running), builds all projects and starts docker compose
up_build: build_mock_api build_cache_service build_api_gateway build_aggregator
	@echo "Stopping docker images (if running...)"
	docker-compose down
	@echo "Building (when required) and starting docker images..."
	docker-compose up --build -d
	@echo "Docker images built and started!"

## down: stop docker compose
down:
	@echo "Stopping docker compose..."
	docker-compose down
	@echo "Done!"

## build_mock_api: builds the mock api binary as a linux executable
build_mock_api:
	@echo "Building mock api binary..."
	cd mock-api-service && env GOOS=linux CGO_ENABLED=0 go build -o ${MOCK_API_BINARY} ./cmd/api
	@echo "Done builing mock api binary!"

## build_mock_api: builds the cache binary as a linux executable
build_cache_service:
	@echo "Building cache binary..."
	cd cache-service && env GOOS=linux CGO_ENABLED=0 go build -o ${CACHE_SERVICE_BINARY} ./cmd/api
	@echo "Done builing cache binary!"

## build_api_gateway: builds api gateway binary as a linux executable
build_api_gateway:
	@echo "Building api gateway binary..."
	cd api-gateway && env GOOS=linux CGO_ENABLED=0 go build -o ${API_GATEWAY_BINARY} ./cmd/api
	@echo "Done builing api gateway binary!"

## build_aggregator: builds aggregator binary as a linux executable
build_aggregator:
	@echo "Building aggregator binary..."
	cd aggregator && env GOOS=linux CGO_ENABLED=0 go build -o ${AGGREGATOR_BINARY} ./cmd/api
	@echo "Done builing aggregator binary!"