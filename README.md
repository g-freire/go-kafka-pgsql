#
### Steps to run
#### 1) Create the network manually using
###### docker network create local-net
#### 2) Then compose up / down the files
###### docker-compose -f "docker-compose.yml" up -d --build
###### docker-compose -f "docker-compose.yml" down
#
### Monitors
##### Grafana ->  localhost:3000 (Import Grafana Postgres Dashboard ID: 9628)
##### Prometheus -> localhost:9090
##### Portainer -> localhost:9000
##### Kowl (Kafka monitor) -> localhost:8080


RUN A CONSUMER
go run .\main.go -t=c

RUN A PRODUCER
go run .\main.go -t=c