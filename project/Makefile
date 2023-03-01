DATA_CONVERTER_BINARY=dataConverterApp
BROKER_BINARY=brokerApp
##builds all services
up_build: build_dataConverter build_broker
	@echo "Stopping docker images"
	docker-compose down
	@echo "Building"
	docker-compose up --build -d
	@echo "Docker imgs built and started"

##builds broker service
build_broker:
	@echo "Building broker binary"
	cd ../broker-service && env GOOS=linux CGO_ENABLED=0 go build -o ${BROKER_BINARY} ./cmd/api
	@echo "Done"
##builds dataConverterService
build_dataConverter:
	@echo "Building dataConverter binary"
	cd ../dataconverter-service && env GOOS=linux CGO_ENABLED=0 go build -o ${DATA_CONVERTER_BINARY} ./cmd/api