# Install an elasticsearch stack

### Operator

Get the actual latest version from elastic.co

First install the CRDs
```shell
$ kubectl create -f https://download.elastic.co/downloads/eck/2.2.0/crds.yaml
```

Install the operator and RBAC rules
```shell
$ kubectl apply -f https://download.elastic.co/downloads/eck/2.2.0/operator.yaml
```
