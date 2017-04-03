# Kubernetes Commands

A set of commands to use for [Kubernetes](https://kubernetes.io/) on both local ([Minikube](https://github.com/kubernetes/minikube)) and production ([Google Container Engine](https://cloud.google.com/container-engine/docs/quickstart)) environments.

## Start Up

```
minikube start
```

## Minikube / Google Container Engine

### Cluster Info

*Minikube*
```
kubectl cluster-info
```

### Dasboard
```
minikube dashboard
```

## Services

### Deploys Image
```
kubectl run redis-master --image=redis --port=6379
```
```
kubectl run app-nginx --image=nginx --port=80
```

### Exposes to Service
```
kubectl expose deployment redis-master --type=NodePort
```

### Lists Services
```
kubectl get services
```
```
kubectl describe services <service-name>
```

### Service Host and Port
*Returns Url*
```
minikube service redis-master --url
```
*Opens Url in Browser*
```
minikube service redis-master
```

## Pods

### Lists Pods
```
kubectl get pods
```
```
kubectl get pods --selector="run=<image-name>" --output=wide
```

### Execument Commands on Pod
```
kubectl exec -it <pod-name> bash
```

*Redis*
```
kubectl exec -t -i <pod-name> redis-cli
```
