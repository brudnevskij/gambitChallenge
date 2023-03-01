# Gambit challenge solution
## Description
My solution introduces micro-service application containing 3 services.
- Broker-service
- DataConverter-service
- Authentication-service
- Postgre server

Application is now configured to run in one minikube kluster, exposing broker-service as a load balancer. During development steps docker-compose was used mainly.

## Service description
### Broker service
Broker service is an API gateway. After receiving a request from client/user, it attempts to athenticate user through authentication-service and then fetches data from dataConverter-service. Broker-service has an exposed 8080 port, that listens for external requests. To get data user must send POST request to the 
>broker-serviceaddress:8080/data

Body of the request must include
>email : email@address
password: password

## Authentication-service
Authentication-service listens for the requests on the non-exposed port 80, on the handler /authenticate. After receiving a request service attempts to validate credentials in postgres user DB. This repo contains a simple SQL querry to populate postgre db with one test user.

## DataConverter-Service 
DataConverter-service listens on non-exposed port 5001, for RPC. After receiving a RPC it fetches data from Gambit API, parses it, converts it to human redable form and serves it to request. There is communication option in http commented.

## Project folder
Project folder contains: build tools implemented in Make, docker-compose.yml to run app without minikube cluster, yml files for k8s deployment.
