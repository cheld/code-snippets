#!/bin/bash

# Runs Heapster in standalone mode
docker run --net=host -d fest/heapster:v1.2.0 \
    --source="kubernetes:http://127.0.0.1:8080?inClusterConfig=false&auth=" \
    --sink="monasca:?user-id=7b23b0d78f124b1aa2439f492dd0e758&password=password&tenant-id=802a70dc02424780adbb7a118e1c1c5c&keystone-url=http://127.0.0.1:5000/v3"

until $(curl --output /dev/null --silent --head --fail http://localhost:8080); do
  printf '.'
  sleep 1
done

docker run --net=host -d gcr.io/google_containers/kubernetes-dashboard-amd64:v1.4.0 --apiserver-host=http://127.0.0.1:8080 \
    --port 9090 --heapster-host=http://127.0.0.1:8082 

docker run --net=host -it -d -v /var/lib/docker/containers:/var/lib/docker/containers:ro -v /var/log/containers:/var/log/containers:ro taimir93/logstash-monasca:v1.1 /run.sh http://127.0.0.1:5607/v3.0 http://127.0.0.1:5000/v3 mini-mon monasca-agent password Default
