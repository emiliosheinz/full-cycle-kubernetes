# Full Cycle K8s

This repository contains the code generated during the Kubernetes module of the Full Cycle course. It contains a simple Go application that prints a message when accessed, and a Dockerfile to build the image which will be used by Kubernetes. Also, inside the `k8s` folder there is the configuration for the deployment of the application, with examples of how to configure the number of replicas, the config maps, the liveness and readiness probes, the horizontal pod autoscaler of the application, and the persistent volume claim. Additionally, there are some other files that can be used on demand, like `statefulset.yaml` and `mysql-service-headless.yaml`, for example, that can be used to create a MySQL database and a Headless service for it.

Kubernetes is an open-source system for automating deployment, scaling, and management of containerized applications. It groups containers that make up an application into logical units for easy management and discovery. Kubernetes follows the the below hierarchical organization of resources:

_Deployments > Replicasets > Pods > Containers_

## 🔧 Running locally

The following steps will guide you through the basic process of running the application locally, using Docker and Kubernetes. But, as mentioned before, this repository contains the configuration for a lot of other things. So, feel free to explore the repository and try to run the application with different configurations.

1. You need to have Kind, Docker, and Kubernetes setup locally
1. Initialize the cluster:
    ```bash
    kind create cluster --config k8s/kind.yaml
    ```
1. Create the service:

    ```bash
    kubectl apply -f k8s/service.yaml
    ```
1. Create the horizontal pod autoscaler:

    ```bash
    kubectl apply -f k8s/hpa.yaml
    ```
1. Create the persistent volume claim:

    ```bash
    kubectl apply -f k8s/pvc.yaml
    ```
1. Start the application:

    ```bash
    kubectl apply -f k8s/deployment.yaml
    ```
1. Expose the service (this is necessary because we are running locally and not in a cloud provider):

    ```bash
    kubectl port-forward service/goserver-service 8000:80
    ```

### 🎯 Expected result

After running the commands above Kubernetes will create the *deployment*, *replicaset* and *3 pods* for the application. To check if the application is running execute the command below:

```
kubectl get pods
```

It should result in something like this:

```
NAME                        READY   STATUS    RESTARTS   AGE
goserver-747c9b6985-jrwf8   1/1     Running   0          19s
goserver-747c9b6985-nwsbn   1/1     Running   0          19s
goserver-747c9b6985-v5bp6   1/1     Running   0          19s
```

Also, when accessing `http://localhost:8000` you should see the message `Hello Full Cycle!!!` printed on the screen.

## 😡 Stressing the application

In order tu stress the application and see the horizontal pod autoscaler in action, you can run the following command:

```bash
kubectl run -it fortio --rm --image=fortio/fortio -- load -qps 1600 -t 220s -c 160 "http://goserver-service/healthz"
```

This will create a pod that will stress the application for 220 seconds, with 1600 queries per second and 160 concurrent connections. To check the current CPU usage and the amount of pod replicas running, you can execute the command below:

```bash
watch -n1 kubectl get hpa
```

## Additional information

### 🧐 Should I use my database inside Kubernetes?

The answer is: it depends. If you have a simple and small application it can be helpful to have it inside Kubernetes, but when talking about a big and critical application usually using a managed database service is the best option. The main reason for that is that the managed database service will take care of the database for you, so you don't need to worry about it. Also, it will be easier to scale the database if necessary, and you will have a lot of other features that will help you to manage your database. 

### 📪 Ingress

Ingress is an API object that manages external access to the services in a cluster, typically HTTP. Ingress may provide load balancing, SSL termination and name-based virtual hosting. Ingress exposes HTTP and HTTPS routes from outside the cluster to services within the cluster. Traffic routing is controlled by rules defined on the Ingress resource.

### 🔒 Cert Manager

Cert-Manager is a Kubernetes add-on to automate the management and issuance of TLS certificates from various issuing sources. It will ensure certificates are valid and up to date periodically, and attempt to renew certificates at an appropriate time before expiry. Currently, this is the most used tool to manage certificates in Kubernetes, and it automatically integrates with Ingress. You can find more information about it [here](https://cert-manager.io/docs/).

### 📦 Namespaces

In Kubernetes, namespaces provides a mechanism for isolating groups of resources within a single cluster. Names of resources need to be unique within a namespace, but not across namespaces. Namespace-based scoping is applicable only for namespaced objects (e.g. Deployments, Services, etc) and not for cluster-wide objects (e.g. StorageClass, Nodes, PersistentVolumes, etc). You can find more information about it [here](https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/).

To create a new namespace you can run the following command:

```bash
kubectl create namespace <namespace-name>
```

To list all namespaces you can run the following command:

```bash
kubectl get namespace
```

### 🔐 Service Accounts and Roles

A service account provides an identity for processes that run in a Pod, and maps to a ServiceAccount object. When you authenticate to the API server, you identify yourself as a particular user. Therefore, Service Accounts along with the proper configuration of Roles can be used to provide an identity for Pods that run in your cluster, that way limiting the access of the Pods to the cluster for example. You can find more information about it [here](https://kubernetes.io/docs/tasks/configure-pod-container/configure-service-account/).

