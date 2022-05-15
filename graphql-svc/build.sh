#!/bin/zsh
GOARCH=amd64 GOOS=linux go build -o app
docker build . -t todos/graphql-svc:latest
kubectl delete -f ./.kube/dev.yaml
kubectl apply -f ./.kube/dev.yaml
