#!/bin/zsh
GOARCH=arm64 GOOS=linux go build -o app
docker build . -t todos/todos-svc:latest
kubectl delete -f ./.kube/dev.yaml
kubectl apply -f ./.kube/dev.yaml
