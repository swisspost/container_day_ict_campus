# Container Day ICT-Campus

Repository containing the resources used at the Container-Day in ICT-Campus.

## Help

- [Docker Cheat Sheet](https://dockerlabs.collabnix.com/docker/cheatsheet/)

## Exercises

### Prepare

You should have:

- cloned the repo
- docker installed
- an editor installed

**Note: for every exercise it's important that you change inside it's directory using `cd`, otherwise the solutions won't work.**

### Beginner

- [Hello World](./beginner/helloworld): Very basic dockerfile that runs a bash script at startup (same as in live demo)
- [Webserver](./beginner/webserver): Very simple webserver that presents a static website

### Advanced

- [Highly available webserver](./advanced/webserver-ha): High-available version of the static website using multiple webservers and a reverse-proxy
- [filebrowser](./advanced/filebrowser/): Basic web-based filebrowser to view/edit/share files on the web
- [Gitea](./advanced/gitea): Lightweight but powerful git server to self-host as container (uses a mysql database)
- [wikijs](./advanced/wikijs/): Wonderful wiki solution for a self-hosted wiki
- [portainer](./advanced/portainer): A Web-Interface for managing containers on docker hosts, kubernetes clusters and more

### Expert

If you want to go deeper into container and kubernetes checkout the following links:

- [Quickstart on how to setup a k3s cluster](https://docs.k3s.io/quick-start): Docs how to setup a lightweight Kubernetes cluster (wait for the K8s intro first)
- [Add more docker hosts to portainer](https://markontech.com/devops/add-a-remote-docker-host-in-portainer/): Scaling using multiple docker hosts managed by a single user-interface
- [How to write a best-practise Dockerfile](https://sysdig.com/blog/dockerfile-best-practices/): To study, those are best-practices everyone should follow but 90% of containers don't (which makes them very insecure)
- [Setup Kubernetes the kubeadm way](https://technat.ch/posts/k8s_kubeadm/): Searching something to do in your project week? This could be your project! Setting up kubernetes using kubeadm is time-intensive but helps you really learn how kubernetes works. Only recommended for people that want to go deep into the topic
