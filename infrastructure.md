# Infrastructure

This page describes what kind of infrastructure is required for the container day.

## Participants

All participants need a VM with Ubuntu Destkop installed on it and the following packages:

- [Docker](https://docs.docker.com/engine/install/ubuntu/)
- [VS Code](https://code.visualstudio.com/docs/setup/linux)

### Access Registry

For the container day we usually create our own registry to avoid any rate-limiting.

Use it like so:

```console
docker login containerday.azurecr.io --username operator -p 234897423894789dsukjfhsdf762d837234df23748
```
