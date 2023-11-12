![278782323-be381844-db4f-412a-ae08-c4de165c3cb7](https://github.com/emiliosheinz/full-cycle-kubernetes/assets/103655828/254f5c94-4146-450f-80f4-df61ce81674d)

# Full Cycle K8s

This repository contains the code generated during the Kubernetes module of the Full Cycle course. It contains a simple Go application that prints a message when accessed, and a Dockerfile to build the image which will be used by Kubernetes. Also, inside the `k8s` folder there is the configuration for the deployment of the application, with examples of how to configure the number of replicas, the config maps, the liveness and readiness probes, and the horizontal pod autoscaler of the application.

Kubernetes is an open-source system for automating deployment, scaling, and management of containerized applications. It groups containers that make up an application into logical units for easy management and discovery. Kubernetes follows the the below hierarchical organization of resources:

_Deployments > Replicasets > Pods > Containers_

## 🔧 Running locally

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

Also, when accessing `http://localhost:8080` you should see the message `Hello Full Cycle!!!` printed on the screen.

## 😡 Stressing the application

In order tu stress the application and see the horizontal pod autoscaler in action, you can run the following command:

```bash
kubectl run -it fortio --rm --image=fortio/fortio -- load -qps 1600 -t 220s -c 160 "http://goserver-service/healthz"
```

This will create a pod that will stress the application for 220 seconds, with 1600 queries per second and 160 concurrent connections. To check the current CPU usage and the amount of pod replicas running, you can execute the command below:

```bash
watch -n1 kubectl get hpa
```bash