#!/bin/bash

# Runs Heapster in standalone mode
docker run --net=host -d gcr.io/google_containers/heapster:v1.1.0 -port 8082 \
    --source="kubernetes:http://127.0.0.1:8080?inClusterConfig=false&auth=" \
    --sink="monasca:?user-id=ee8b70289cbb45d88e4befb17a0e6c55&password=password&monasca-url=http://127.0.0.1:8070/v2.0&keystone-url=http://127.0.0.1:5000/v3"

sleep 5

docker run --net=host -d gcr.io/google_containers/kubernetes-dashboard-amd64:v1.4.0 --apiserver-host=http://127.0.0.1:8080 \
    --port 9090 --heapster-host=http://127.0.0.1:8082 
