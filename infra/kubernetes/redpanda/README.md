# Installing the redpanda-operator

### Get the latest version from github
```shell
$ export VERSION=$(curl -s https://api.github.com/repos/redpanda-data/redpanda/releases/latest | jq -r .tag_name)
```

### Install the CRDs
```shell
$ noglob kubectl apply -k https://github.com/redpanda-data/redpanda/src/go/k8s/config/crd?ref=$VERSION
```

### Install helm and the repo
If not already installed, run the following to install helm and the repo
```shell
$ brew install helm
$ helm repo add redpanda https://charts.vectorized.io/ && \
```

### Install the operator
```shell
$ helm repo update
$ helm install \
  redpanda-operator \
  redpanda/redpanda-operator \
  --namespace redpanda-system \
  --create-namespace \
  --version $VERSION
```
