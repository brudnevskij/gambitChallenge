DATA_CONVERTER_BINARY=dataConverterApp
BROKER_BINARY=brokerApp
AUTHENTICATION_BINARY=authApp
##builds all services
up_build: build_dataConverter build_broker build_authentication
	@echo "Stopping docker images"
	docker-compose down
	@echo "Building"
	docker-compose up --build -d
	@echo "Docker imgs built and started"
##build all binaries
build_all: build_dataConverter build_broker build_authentication
##builds broker service
build_broker:
	@echo "Building broker binary"
	cd ../broker-service && env GOOS=linux CGO_ENABLED=0 go build -o ${BROKER_BINARY} ./cmd/api
	@echo "Done"
##builds dataConverterService
build_dataConverter:
	@echo "Building dataConverter binary"
	cd ../dataconverter-service && env GOOS=linux CGO_ENABLED=0 go build -o ${DATA_CONVERTER_BINARY} ./cmd/api
	@echo "Done"
##builds auth server
build_authentication:
	@echo "Building authentication binary"
	cd ../authentication-service && env GOOS=linux CGO_ENABLED=0 go build -o ${AUTHENTICATION_BINARY} ./cmd/api
	@echo "Done"
start_local_postgres:
	@echo "Starting local postgres"
	docker-compose -f postgres.yml up -d
	@echo "Done"
##start and deploy k8s pods
deploy_and_expose: start_local_postgres
	@echo "Deploying k8s pods and exposing broker on port 8080..."
	kubectl apply -f k8s
	kubectl delete svc broker-service
	kubectl expose deployment broker-service --type=LoadBalancer --port=8080 --target-port=8080
	@echo "Done! Broker is now pending for expernal ip"
