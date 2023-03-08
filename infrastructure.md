# Infrastructure

This page describes what kind of infrastructure is required for the container day.

## Participants

All participants need a VM with Ubuntu Destkop installed on it and the following packages:

- [Docker](https://docs.docker.com/engine/install/ubuntu/)
- [Kubectl](https://kubernetes.io/docs/tasks/tools/install-kubectl-linux/)

## Kubernetes Cluster

For the Kubernetes-Part a K8s cluster is needed. For simplicity this one is created in Azure Cloud with the following options:

| Key                             | Value                                                       |
|---------------------------------|-------------------------------------------------------------|
| Subscription                    | Azure for Students                                          |
| Resource Group                  | containerday                                                |
| Cluster preset configuration    | Dev/Test                                                    |
| Cluster Name                    | containerday                                                |
| Region                          | Switzerland North                                           |
| Availability Zones              | 1,2,3                                                       |
| K8s Version                     | 1.25.4                                                      |
| AKS pricing tier                | free                                                        |
| Node size                       | B2s (2/4)                                                   |
| Node count range                | 1-2                                                         |
| Authentization & Authorization  | Local accounts with Kubernetes RBAC                         |
| Enable HTTP application routing | true                                                        |
| Azure Container Registry        | Create new: name: containerday enable admin user SKU: basic |
| Enable recommended alert rules  | disabled                                                    |
| Edit Infrastructure RG name     | containerday_infra                                          |

### Access Cluster

Once the cluster is built, you can access it like so:

```console
az login # Open a browser, select your account
az account list # Get the ID of your subscription
az account set --subscription <ID> # Set the current active subscription
az aks get-credentials --admin --name containerday --resource-group containerday # Saves the kubeconfig into ~/.kube/config
```

Note: the `--admin` is important so that the kubeconfig doesn't rely on [kubelogin](https://github.com/Azure/kubelogin) to be authenticated. This way the kubeconfig file can be shared with anyone and he is able to access the cluster.

### Access Registry

The container registry that was created too with the wizard can be found in the portal. Under it's menu "Access keys" you can find two passwords you can use to log into the registry.

Use them like so:

```console
docker login containerday.azurecr.io --username containerday -p 234897423894789dsukjfhsdf762d837234df23748
```
