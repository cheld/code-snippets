# Runs Heapster in standalone mode
docker run --net=host -d gcr.io/google_containers/heapster:v1.2.0 -port 8082 \
    --source=kubernetes:http://127.0.0.1:8080?inClusterConfig=false&auth=""

sleep 5

docker run --net=host -d gcr.io/google_containers/kubernetes-dashboard-amd64:v1.4.0 --apiserver-host=http://127.0.0.1:8080 \
    --port 9090 --heapster-host=http://127.0.0.1:8082 
